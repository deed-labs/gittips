package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/messages"
	"github.com/deed-labs/gittips/bot/internal/parser"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type GitHubService struct {
	client *github.Client

	owners   Owners
	bounties Bounties
	commands Commands
}

func NewGitHubService(httpClient *http.Client, owners Owners, bounties Bounties, commands Commands) *GitHubService {
	client := github.NewClient(httpClient)

	return &GitHubService{
		client:   client,
		owners:   owners,
		bounties: bounties,
		commands: commands,
	}
}

func (s *GitHubService) ProcessOrganizationInstallation(ctx context.Context, payload ghHooks.InstallationPayload) error {
	return s.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (s *GitHubService) ProcessRepositoriesInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error {
	return s.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (s *GitHubService) processInstallation(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error {
	ownerExists, err := s.owners.Exists(ctx, id)
	if err != nil {
		return err
	}

	var name string
	switch ownerType {
	case "User":
		c, err := s.getUserClient(ctx, login)
		if err != nil {
			return fmt.Errorf("get user client: %w", err)
		}
		user, _, err := c.Users.Get(ctx, login)
		if err != nil {
			return fmt.Errorf("get user info: %w", err)
		}
		name = *user.Name

	case "Organization":
		c, err := s.getOrganizationClient(ctx, login)
		if err != nil {
			return fmt.Errorf("get organization client: %w", err)
		}
		org, _, err := c.Organizations.Get(ctx, login)
		if err != nil {
			return fmt.Errorf("get organization info: %w", err)
		}
		name = *org.Name
	}

	if !ownerExists {
		err := s.owners.Create(ctx, id, login, name, url, avatarURL, ownerType)
		if err != nil {
			return fmt.Errorf("create owner: %w", err)
		}
	}

	return nil
}

func (s *GitHubService) ProcessIssueEvent(ctx context.Context, payload ghHooks.IssuesPayload) error {
	isMember, client, err := s.getClientAndMembership(ctx, payload.Sender.Login, payload.Repository.Owner.Login,
		payload.Repository.Owner.Type)
	if err != nil {
		return err
	}

	labelNames := make([]string, 0, len(payload.Issue.Labels))
	for _, v := range payload.Issue.Labels {
		labelNames = append(labelNames, v.Name)
	}

	switch {
	case parser.SearchLabel(parser.CreateBountyLabel, labelNames):
		if payload.Action == "opened" {
			if !isMember {
				reply := fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.NotEnoughPermissionsToCreateBounty)
				comment := &github.IssueComment{
					Body: &reply,
				}

				_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.Issue.Number), comment)
				if err != nil {
					return fmt.Errorf("create comment: %w", err)
				}

				_, err = client.Issues.RemoveLabelForIssue(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.Issue.Number), string(parser.CreateBountyLabel))
				if err != nil {
					return fmt.Errorf("delete label: %w", err)
				}

				return nil
			}

			err := s.bounties.Create(ctx, payload.Issue.ID, payload.Repository.Owner.ID,
				payload.Issue.Title, payload.Issue.URL, payload.Issue.Body)
			if err != nil {
				return fmt.Errorf("create bounty: %w", err)
			}
		} else if payload.Action == "closed" {
			err := s.bounties.Close(ctx, payload.Issue.ID)
			if err != nil {
				return fmt.Errorf("close bounty: %w", err)
			}
		} else if payload.Action == "deleted" {
			err := s.bounties.Delete(ctx, payload.Issue.ID)
			if err != nil {
				return fmt.Errorf("delete bounty: %w", err)
			}
		}
	}

	return nil
}

func (s *GitHubService) ProcessIssueComment(ctx context.Context, payload ghHooks.IssueCommentPayload) error {
	parsedBody := parser.Parse(payload.Comment.Body)
	if len(parsedBody.Commands) == 0 {
		// Ignore comments without any commands.
		return nil
	}

	isMember, client, err := s.getClientAndMembership(ctx, payload.Sender.Login, payload.Repository.Owner.Login,
		payload.Repository.Owner.Type)
	if err != nil {
		return err
	}

	if !isMember {
		reply := fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.NotEnoughPermissionsToRunCommands)
		comment := &github.IssueComment{
			Body: &reply,
		}

		_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
			payload.Repository.Name, int(payload.Issue.Number), comment)
		if err != nil {
			return fmt.Errorf("create comment: %w", err)
		}
	}

	for _, v := range parsedBody.Commands {
		cmd := s.commands.Parse(v)
		if cmd == nil {
			continue
		}

		switch c := cmd.(type) {
		case *SendPaymentCommand:
			toOwner, _, err := client.Users.Get(ctx, c.To)
			if err != nil {
				reply := fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.InvalidUserInput)
				comment := &github.IssueComment{
					Body: &reply,
				}

				_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.Issue.Number), comment)
				if err != nil {
					return fmt.Errorf("create comment: %w", err)
				}

				return nil
			}
			if err := c.Run(ctx, *toOwner.ID, c.Value); err != nil {
				var reply string
				switch {
				case errors.Is(err, ErrInvalidValue):
					reply = fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.InvalidValueInput)
				case errors.Is(err, ErrUserNotFound):
					msg := fmt.Sprintf(messages.UserHasNoWalletTmpl, c.To)
					reply = fmt.Sprintf("@%s\n%s", payload.Sender.Login, msg)
				default:
					return fmt.Errorf("run command: %w", err)
				}

				comment := &github.IssueComment{
					Body: &reply,
				}

				_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.Issue.Number), comment)
				if err != nil {
					return fmt.Errorf("create comment: %w", err)
				}
			}

			msg := fmt.Sprintf(messages.PaymentSentTmpl, c.To)
			comment := &github.IssueComment{
				Body: &msg,
			}

			_, _, err = client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
				payload.Repository.Name, int(payload.Issue.Number), comment)
			if err != nil {
				return fmt.Errorf("create comment: %w", err)
			}

			return nil
		case *SetWalletCommand:
			if err := c.Run(ctx, payload.Sender.ID, c.WalletAddress); err != nil {
				return fmt.Errorf("run command: %w", err)
			}
		case *SetRewardCommand:
			if err := c.Run(ctx, payload.Issue.ID, c.RewardValue); err != nil {
				return fmt.Errorf("run command: %w", err)
			}
		case *CloseCommand:
			newState := "closed"
			issue := &github.IssueRequest{
				State: &newState,
			}
			_, _, err := client.Issues.Edit(ctx, payload.Repository.Owner.Login,
				payload.Repository.Name, int(payload.Issue.Number), issue)
			if err != nil {
				return fmt.Errorf("edit issue: %w", err)
			}
		default:
			continue
		}
	}

	return nil
}

func (s *GitHubService) ProcessPRComment(ctx context.Context, payload ghHooks.PullRequestReviewCommentPayload) error {
	parsedBody := parser.Parse(payload.Comment.Body)
	if len(parsedBody.Commands) == 0 {
		// Ignore comments without any commands.
		return nil
	}

	isMember, client, err := s.getClientAndMembership(ctx, payload.Sender.Login, payload.Repository.Owner.Login,
		payload.Repository.Owner.Type)
	if err != nil {
		return err
	}

	if !isMember {
		reply := fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.NotEnoughPermissionsToRunCommands)
		comment := &github.PullRequestComment{
			Body: &reply,
		}

		_, _, err := client.PullRequests.CreateComment(ctx, payload.Repository.Owner.Login,
			payload.Repository.Name, int(payload.PullRequest.Number), comment)
		if err != nil {
			return fmt.Errorf("create comment: %w", err)
		}
	}

	for _, v := range parsedBody.Commands {
		cmd := s.commands.Parse(v)
		if cmd == nil {
			continue
		}

		switch c := cmd.(type) {
		case *SendPaymentCommand:
			toOwner, _, err := client.Users.Get(ctx, c.To)
			if err != nil {
				reply := fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.InvalidUserInput)
				comment := &github.IssueComment{
					Body: &reply,
				}

				_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.PullRequest.Number), comment)
				if err != nil {
					return fmt.Errorf("create comment: %w", err)
				}

				return nil
			}

			if err := c.Run(ctx, *toOwner.ID, c.Value); err != nil {
				var reply string
				switch {
				case errors.Is(err, ErrInvalidValue):
					reply = fmt.Sprintf("@%s\n%s", payload.Sender.Login, messages.InvalidValueInput)
				case errors.Is(err, ErrUserNotFound):
					msg := fmt.Sprintf(messages.UserHasNoWalletTmpl, c.To)
					reply = fmt.Sprintf("@%s\n%s", payload.Sender.Login, msg)
				default:
					return fmt.Errorf("run command: %w", err)
				}

				comment := &github.IssueComment{
					Body: &reply,
				}

				_, _, err := client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
					payload.Repository.Name, int(payload.PullRequest.Number), comment)
				if err != nil {
					return fmt.Errorf("create comment: %w", err)
				}
			}

			msg := fmt.Sprintf(messages.PaymentSentTmpl, c.To)
			comment := &github.IssueComment{
				Body: &msg,
			}

			_, _, err = client.Issues.CreateComment(ctx, payload.Repository.Owner.Login,
				payload.Repository.Name, int(payload.PullRequest.Number), comment)
			if err != nil {
				return fmt.Errorf("create comment: %w", err)
			}
		case *SetWalletCommand:
			if err := c.Run(ctx, payload.Sender.ID, c.WalletAddress); err != nil {
				return fmt.Errorf("run command: %w", err)
			}
		case *SetRewardCommand:
			// Not supported for pull request for now
			continue
		case *CloseCommand:
			newState := "closed"
			pr := &github.PullRequest{
				State: &newState,
			}
			_, _, err := client.PullRequests.Edit(ctx, payload.Repository.Owner.Login,
				payload.Repository.Name, int(payload.PullRequest.Number), pr)
			if err != nil {
				return fmt.Errorf("edit issue: %w", err)
			}
		default:
			continue
		}
	}

	return nil
}

func (s *GitHubService) ProcessInstallationSetup(ctx context.Context, installationId int64, walletAddress string) error {
	installation, _, err := s.client.Apps.GetInstallation(ctx, installationId)
	if err != nil {
		return fmt.Errorf("get installation: %w", err)
	}

	return s.owners.LinkWithWallet(ctx, *installation.Account.ID, walletAddress)
}

func (s *GitHubService) getClientAndMembership(ctx context.Context, user string, owner string, ownerType string) (bool, *github.Client, error) {
	var (
		isMember bool
		client   *github.Client
	)
	switch ownerType {
	case "User":
		c, err := s.getUserClient(ctx, user)
		if err != nil {
			return false, nil, fmt.Errorf("get user client: %w", err)
		}
		client = c
		// NOTE: Only repository owner is allowed to create bounties.
		isMember = user == owner
	case "Organization":
		c, err := s.getOrganizationClient(ctx, owner)
		if err != nil {
			return false, nil, fmt.Errorf("get organization client: %w", err)
		}
		client = c

		msStatus, _, err := client.Organizations.IsMember(ctx, owner, user)
		if err != nil {
			return false, nil, fmt.Errorf("get membership status: %w", err)
		}
		isMember = msStatus
	}

	return isMember, client, nil
}

func (s *GitHubService) getUserClient(ctx context.Context, user string) (*github.Client, error) {
	installation, _, err := s.client.Apps.FindUserInstallation(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("find user installation: %w", err)
	}

	// TODO: store client to cache

	return s.getClient(ctx, *installation.ID)
}

func (s *GitHubService) getOrganizationClient(ctx context.Context, org string) (*github.Client, error) {
	installation, _, err := s.client.Apps.FindOrganizationInstallation(ctx, org)
	if err != nil {
		return nil, fmt.Errorf("find organization installation: %w", err)
	}

	// TODO: store client to cache

	return s.getClient(ctx, *installation.ID)

}
func (s *GitHubService) getClient(ctx context.Context, installationId int64) (*github.Client, error) {
	token, _, err := s.client.Apps.CreateInstallationToken(ctx, installationId, &github.InstallationTokenOptions{})
	if err != nil {
		return nil, fmt.Errorf("create installation token: %w", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.GetToken()},
	)
	oAuthClient := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(oAuthClient)

	return client, nil
}

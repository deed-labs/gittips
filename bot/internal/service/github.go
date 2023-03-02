package service

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/messages"
	"github.com/deed-labs/gittips/bot/pkg/parser"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"github.com/google/go-github/v50/github"
)

type GithubService struct {
	client *github.Client

	owners   Owners
	bounties Bounties
}

func NewGithubService(httpClient *http.Client, owners Owners, bounties Bounties) *GithubService {
	client := github.NewClient(httpClient)

	return &GithubService{
		client:   client,
		owners:   owners,
		bounties: bounties,
	}
}

func (s *GithubService) ProcessOrganizationInstallation(ctx context.Context, payload ghHooks.InstallationPayload) error {
	return s.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (s *GithubService) ProcessRepositoriesInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error {
	return s.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (s *GithubService) processInstallation(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error {
	ownerExists, err := s.owners.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !ownerExists {
		err := s.owners.Create(ctx, id, login, url, avatarURL, ownerType)
		if err != nil {
			return fmt.Errorf("create owner: %w", err)
		}
	}

	return nil
}

func (s *GithubService) ProcessIssueEvent(ctx context.Context, payload ghHooks.IssuesPayload) error {
	var (
		isMember bool
		client   *github.Client
	)
	switch payload.Repository.Owner.Type {
	case "User":
		c, err := s.getUserClient(ctx, payload.Repository.Owner.Login)
		if err != nil {
			return err
		}
		client = c
		// NOTE: Only repository owner is allowed to create bounties.
		isMember = payload.Sender.Login == payload.Repository.Owner.Login
	case "Organization":
		c, err := s.getOrganizationClient(ctx, payload.Repository.Owner.Login)
		if err != nil {
			return err
		}
		client = c

		msStatus, _, err := client.Organizations.IsMember(ctx, payload.Repository.Owner.Login, payload.Sender.Login)
		if err != nil {
			return fmt.Errorf("get membership status: %w", err)
		}
		isMember = msStatus

	}

	labelNames := make([]string, 0, len(payload.Issue.Labels))
	for _, v := range payload.Issue.Labels {
		labelNames = append(labelNames, v.Name)
	}

	switch {
	case parser.SearchLabel(parser.CreateBountyLabel, labelNames):
		if payload.Action == "opened" {
			if !isMember {
				comment := &github.IssueComment{
					Body: &messages.NotEnoughPermissionsToCreateBounty,
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
		} else if payload.Action == "closed" || payload.Action == "deleted" {
			err := s.bounties.Delete(ctx, payload.Issue.ID)
			if err != nil {
				return fmt.Errorf("delete bounty: %w", err)
			}
		}
	}

	return nil
}

func (s *GithubService) ProcessIssueComment(ctx context.Context, payload ghHooks.IssueCommentPayload) error {
	return nil
}

func (s *GithubService) ProcessNewPR(ctx context.Context, payload ghHooks.PullRequestPayload) error {
	return nil
}

func (s *GithubService) ProcessPRComment(ctx context.Context, payload ghHooks.PullRequestReviewCommentPayload) error {
	return nil
}

func (s *GithubService) getUserClient(ctx context.Context, user string) (*github.Client, error) {
	installation, _, err := s.client.Apps.FindUserInstallation(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("find user installation: %w", err)
	}

	// TODO: store client to cache

	return s.getClient(ctx, *installation.ID)
}

func (s *GithubService) getOrganizationClient(ctx context.Context, org string) (*github.Client, error) {
	installation, _, err := s.client.Apps.FindOrganizationInstallation(ctx, org)
	if err != nil {
		return nil, fmt.Errorf("find organization installation: %w", err)
	}

	// TODO: store client to cache

	return s.getClient(ctx, *installation.ID)

}
func (s *GithubService) getClient(ctx context.Context, installationId int64) (*github.Client, error) {
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

package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/parser"
	"github.com/deed-labs/gittips/bot/internal/service"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"github.com/google/go-github/v50/github"
)

type GitHub struct {
	secret   string
	client   *github.Client
	services *service.Services
}

func New(secret string, httpClient *http.Client, services *service.Services) *GitHub {
	client := github.NewClient(httpClient)

	return &GitHub{
		secret:   secret,
		client:   client,
		services: services,
	}
}

func (gh *GitHub) processOrganizationInstallation(ctx context.Context, payload ghHooks.InstallationPayload) error {
	return gh.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (gh *GitHub) processRepositoriesInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error {
	return gh.processInstallation(
		ctx,
		payload.Installation.Account.ID,
		payload.Installation.Account.Login,
		payload.Installation.Account.URL,
		payload.Installation.Account.AvatarURL,
		payload.Installation.Account.Type,
	)
}

func (gh *GitHub) processInstallation(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error {
	ownerExists, err := gh.services.Owners.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !ownerExists {
		err := gh.services.Owners.Create(ctx, id, login, url, avatarURL, ownerType)
		if err != nil {
			return fmt.Errorf("create owner: %w", err)
		}
	}

	return nil
}

func (gh *GitHub) processIssueEvent(ctx context.Context, payload ghHooks.IssuesPayload) error {
	// TODO: check user permissions

	labelNames := make([]string, 0, len(payload.Issue.Labels))
	for _, v := range payload.Issue.Labels {
		labelNames = append(labelNames, v.Name)
	}

	switch {
	case parser.SearchLabel(parser.CreateBountyLabel, labelNames):
		if payload.Action == "opened" {
			err := gh.services.Bounties.Create(ctx, payload.Issue.ID, payload.Repository.Owner.ID,
				payload.Issue.Title, payload.Issue.URL, payload.Issue.Body)
			if err != nil {
				return fmt.Errorf("create bounty: %w", err)
			}
		} else if payload.Action == "closed" || payload.Action == "deleted" {
			err := gh.services.Bounties.Delete(ctx, payload.Issue.ID)
			if err != nil {
				return fmt.Errorf("delete bounty: %w", err)
			}
		}
	}

	return nil
}

func (gh *GitHub) processIssueComment(ctx context.Context, payload ghHooks.IssueCommentPayload) error {
	return nil
}

func (gh *GitHub) processNewPR(ctx context.Context, payload ghHooks.PullRequestPayload) error {
	return nil
}

func (gh *GitHub) processPRComment(ctx context.Context, payload ghHooks.PullRequestReviewCommentPayload) error {
	return nil
}

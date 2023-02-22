package github

import (
	"context"
	"fmt"
	"github.com/deed-labs/openroll/bot/internal/parser"
	"github.com/deed-labs/openroll/bot/internal/service"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"net/http"

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

func (gh *GitHub) Handler() (*Handler, error) {
	hook, err := ghHooks.New(ghHooks.Options.Secret(gh.secret))
	if err != nil {
		return nil, fmt.Errorf("create hook: %w", err)
	}

	h := &Handler{
		github: gh,
		hook:   hook,
	}

	return h, nil
}

func (gh *GitHub) processInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error {
	ownerExists, err := gh.services.Owners.Exists(ctx, payload.Installation.Account.ID)
	if err != nil {
		return err
	}

	if !ownerExists {
		err := gh.services.Owners.Create(
			ctx,
			payload.Installation.Account.ID,
			payload.Installation.Account.Login,
			payload.Installation.Account.URL,
			payload.Installation.Account.AvatarURL,
		)
		if err != nil {
			return fmt.Errorf("create owner: %w", err)
		}
	}

	return nil
}

func (gh *GitHub) processNewIssue(ctx context.Context, payload ghHooks.IssuesPayload) error {
	// TODO: check user permissions

	labelNames := make([]string, 0, len(payload.Issue.Labels))
	for _, v := range payload.Issue.Labels {
		labelNames = append(labelNames, v.Name)
	}

	switch {
	case parser.SearchLabel(parser.CreateBountyLabel, labelNames):
		err := gh.services.Bounties.Create(ctx, payload.Repository.Owner.ID, payload.Issue.Title,
			payload.Issue.URL, payload.Issue.Body)
		if err != nil {
			return fmt.Errorf("create bounty: %w", err)
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

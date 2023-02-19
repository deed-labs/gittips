package github

import (
	ghHooks "github.com/go-playground/webhooks/v6/github"

	"github.com/google/go-github/v50/github"
)

type GitHub struct {
	client *github.Client
}

func New(client *github.Client) *GitHub {
	return &GitHub{client: client}
}

func (gh *GitHub) processIssueComment(payload ghHooks.IssueCommentPayload) error {
	return nil
}

func (gh *GitHub) processPRComment(payload ghHooks.PullRequestReviewCommentPayload) error {
	return nil
}

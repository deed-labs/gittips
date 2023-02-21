package github

import (
	"errors"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"net/http"
)

type Handler struct {
	github *GitHub
	hook   *ghHooks.Webhook
}

var eventsForHandling = []ghHooks.Event{
	ghHooks.InstallationRepositoriesEvent,
	ghHooks.IssuesEvent,
	ghHooks.IssueCommentEvent,
	ghHooks.PullRequestEvent,
	ghHooks.PullRequestReviewCommentEvent,
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	payload, err := h.hook.Parse(r, eventsForHandling...)
	if err != nil {
		if errors.Is(err, ghHooks.ErrEventNotFound) {
			return
		}

		// TODO: handle error
	}

	ctx := r.Context()

	switch p := payload.(type) {
	case ghHooks.InstallationRepositoriesPayload:
		err = h.github.processInstallation(ctx, p)
	case ghHooks.IssuesPayload:
		err = h.github.processNewIssue(ctx, p)
	case ghHooks.IssueCommentPayload:
		err = h.github.processIssueComment(ctx, p)
	case ghHooks.PullRequestPayload:
		err = h.github.processNewPR(ctx, p)
	case ghHooks.PullRequestReviewCommentPayload:
		err = h.github.processPRComment(ctx, p)
	default:
		// unknown payload
		return
	}

	if err != nil {
		// TODO: handle error
	}
}

package github

import (
	"errors"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
)

type Handler struct {
	github *GitHub
	hook   *github.Webhook
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	payload, err := h.hook.Parse(r, github.IssueCommentEvent, github.PullRequestReviewCommentEvent)
	if err != nil {
		if errors.Is(err, github.ErrEventNotFound) {
			return
		}

		// TODO: handle error
	}

	switch p := payload.(type) {
	case github.IssueCommentPayload:
		err = h.github.processIssueComment(p)
	case github.PullRequestReviewCommentPayload:
		err = h.github.processPRComment(p)
	default:
		// unknown payload
		return
	}

	if err != nil {
		// TODO: handle error
	}
}

package github

import (
	"errors"
	"fmt"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	github *GitHub
	hook   *ghHooks.Webhook
	logger *zap.SugaredLogger
}

var eventsForHandling = []ghHooks.Event{
	ghHooks.InstallationEvent,
	ghHooks.InstallationRepositoriesEvent,
	ghHooks.IssuesEvent,
	ghHooks.IssueCommentEvent,
	ghHooks.PullRequestEvent,
	ghHooks.PullRequestReviewCommentEvent,
}

func NewHandler(github *GitHub, logger *zap.SugaredLogger) (*Handler, error) {
	hook, err := ghHooks.New(ghHooks.Options.Secret(github.secret))
	if err != nil {
		return nil, fmt.Errorf("create hook: %w", err)
	}

	h := &Handler{
		github: github,
		hook:   hook,
		logger: logger,
	}

	return h, nil
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	payload, err := h.hook.Parse(r, eventsForHandling...)
	if err != nil {
		if errors.Is(err, ghHooks.ErrEventNotFound) {
			w.WriteHeader(http.StatusOK)

			return
		}

		h.logger.Errorf("failed to parse hook: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	ctx := r.Context()

	h.logger.Info("new webhook handled")

	switch p := payload.(type) {
	case ghHooks.InstallationPayload:
		err = h.github.processOrganizationInstallation(ctx, p)
	case ghHooks.InstallationRepositoriesPayload:
		err = h.github.processRepositoriesInstallation(ctx, p)
	case ghHooks.IssuesPayload:
		err = h.github.processNewIssue(ctx, p)
	case ghHooks.IssueCommentPayload:
		err = h.github.processIssueComment(ctx, p)
	case ghHooks.PullRequestPayload:
		err = h.github.processNewPR(ctx, p)
	case ghHooks.PullRequestReviewCommentPayload:
		err = h.github.processPRComment(ctx, p)
	default:
		h.logger.Warn("unknown webhook handled")

		return
	}

	if err != nil {
		h.logger.Errorf("failed to process hook: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

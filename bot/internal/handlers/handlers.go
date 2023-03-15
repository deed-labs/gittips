package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deed-labs/gittips/bot/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type Handlers struct {
	hook     *ghHooks.Webhook
	services *service.Services
	http     *chi.Mux
	logger   *zap.SugaredLogger
}

var githubEvents = []ghHooks.Event{
	ghHooks.InstallationEvent,
	ghHooks.InstallationRepositoriesEvent,
	ghHooks.IssuesEvent,
	ghHooks.IssueCommentEvent,
	ghHooks.PullRequestReviewCommentEvent,
}

func New(services *service.Services, whSecret string, logger *zap.SugaredLogger) (*Handlers, error) {
	hook, err := ghHooks.New(ghHooks.Options.Secret(whSecret))
	if err != nil {
		return nil, fmt.Errorf("create hook: %w", err)
	}

	h := &Handlers{
		hook:     hook,
		services: services,
		logger:   logger,
	}

	corsCfg := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"Accept", "Content-Type", "Accept-Encoding"},
	})

	r := chi.NewRouter()
	r.Use(corsCfg.Handler)
	r.Use(middleware.DefaultLogger)
	r.Post("/setup", h.handleSetup)
	r.Post("/github", h.handleGitHubWebhook)
	r.Get("/api/bounties", h.handleGetBounties)
	r.Get("/api/installation/{address}", h.handleGetInstallation)
	r.Get("/api/owner/{id}", h.handleGetOwner)

	h.http = r

	return h, nil
}

func (h *Handlers) HTTP() http.Handler {
	return h.http
}

func (h *Handlers) handleGitHubWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := h.hook.Parse(r, githubEvents...)
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
		err = h.services.GitHub.ProcessOrganizationInstallation(ctx, p)
	case ghHooks.InstallationRepositoriesPayload:
		err = h.services.GitHub.ProcessRepositoriesInstallation(ctx, p)
	case ghHooks.IssuesPayload:
		err = h.services.GitHub.ProcessIssueEvent(ctx, p)
	case ghHooks.IssueCommentPayload:
		err = h.services.GitHub.ProcessIssueComment(ctx, p)
	case ghHooks.PullRequestReviewCommentPayload:
		err = h.services.GitHub.ProcessPRComment(ctx, p)
	default:
		h.logger.Warn("unknown webhook handled")

		return
	}

	if err != nil {
		h.logger.Errorf("failed to process hook: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) handleGetBounties(w http.ResponseWriter, r *http.Request) {
	bounties, err := h.services.Bounties.GetAll(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	bountiesList := make([]BountyWithOwner, 0, len(bounties))
	for _, v := range bounties {
		bountiesList = append(bountiesList, BountyWithOwner{
			OwnerID:        v.OwnerID,
			Owner:          v.OwnerLogin,
			OwnerURL:       v.OwnerURL,
			OwnerAvatarURL: v.OwnerAvatarURL,
			OwnerType:      v.OwnerType,
			Title:          v.Title,
			URL:            v.URL,
			Reward:         v.Reward.String(),
		})
	}

	resp := BountyResponse{Bounties: bountiesList}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (h *Handlers) handleSetup(w http.ResponseWriter, r *http.Request) {
	var req SetupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)

		return
	}

	if err := h.services.GitHub.ProcessInstallationSetup(r.Context(), req.InstallationID, req.WalletAddress); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) handleGetInstallation(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	installation, err := h.services.Owners.GetInstallationInfo(r.Context(), address)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	resp := InstallationInfoResponse{
		Installed: installation.Installed,
		OwnerName: installation.OwnerName,
		OwnerID:   installation.OwnerID,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (h *Handlers) handleGetOwner(w http.ResponseWriter, r *http.Request) {
	ownerIdParam := chi.URLParam(r, "id")

	ownerId, err := strconv.Atoi(ownerIdParam)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	owner, err := h.services.Owners.Get(r.Context(), int64(ownerId))
	if err != nil && errors.Is(err, service.ErrOwnerNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}

	resp := OwnerInfoResponse{
		Name:              owner.Owner.Name,
		AvatarURL:         owner.Owner.AvatarURL,
		TotalBudget:       owner.TotalBudget.String(),
		AvailableBudget:   owner.AvailableBudget.String(),
		TotalBounties:     owner.TotalBounties,
		AvailableBounties: owner.AvailableBounties,
		Bounties:          make([]Bounty, 0, len(owner.Bounties)),
	}
	for _, bounty := range owner.Bounties {
		resp.Bounties = append(resp.Bounties, BountyFromEntity(bounty))
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

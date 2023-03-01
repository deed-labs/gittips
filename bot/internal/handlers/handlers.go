package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	services *service.Services
	http     *chi.Mux
}

func New(ghHandler http.HandlerFunc, services *service.Services) *Handlers {
	r := chi.NewRouter()

	h := new(Handlers)
	h.services = services

	r.Use(middleware.DefaultLogger)
	r.Post("/github", ghHandler)
	r.Get("/api/bounties", h.handleGetBounties)

	h.http = r

	return h
}

func (h *Handlers) HTTP() http.Handler {
	return h.http
}

func (h *Handlers) handleGetBounties(w http.ResponseWriter, r *http.Request) {
	bounties, err := h.services.Bounties.GetAll(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	bountiesList := make([]Bounty, 0, len(bounties))
	for _, v := range bounties {
		bountiesList = append(bountiesList, Bounty{
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

	bountiesResponse := BountyResponse{Bounties: bountiesList}

	if err := json.NewEncoder(w).Encode(bountiesResponse); err != nil {

	}
}

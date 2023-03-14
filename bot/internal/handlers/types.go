package handlers

import "github.com/deed-labs/gittips/bot/internal/entity"

type BountyResponse struct {
	Bounties []BountyWithOwner `json:"bounties"`
}

type Bounty struct {
	OwnerID int64  `json:"owner_id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Reward  string `json:"reward"`
}

func BountyFromEntity(e *entity.Bounty) Bounty {
	return Bounty{
		OwnerID: e.OwnerID,
		Title:   e.Title,
		URL:     e.URL,
		Reward:  e.Reward.String(),
	}
}

type BountyWithOwner struct {
	OwnerID        int64  `json:"owner_id"`
	Owner          string `json:"owner"`
	OwnerURL       string `json:"owner_url"`
	OwnerAvatarURL string `json:"owner_avatar_url"`
	OwnerType      string `json:"owner_type"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Reward         string `json:"reward"`
}

func BountyWithOwnerFromEntity(e *entity.BountyWithOwner) BountyWithOwner {
	return BountyWithOwner{
		OwnerID:        e.OwnerID,
		Owner:          e.OwnerLogin,
		OwnerURL:       e.OwnerURL,
		OwnerAvatarURL: e.OwnerAvatarURL,
		OwnerType:      e.OwnerType,
		Title:          e.Title,
		URL:            e.URL,
		Reward:         e.Reward.String(),
	}
}

type InstallationInfoResponse struct {
	Installed bool   `json:"installed"`
	OwnerName string `json:"owner_name"`
	OwnerID   int64  `json:"owner_id"`
}

type OwnerInfoResponse struct {
	Name              string   `json:"name"`
	AvatarURL         string   `json:"avatar_url"`
	TotalBudget       string   `json:"total_budget"`
	AvailableBudget   string   `json:"available_budget"`
	TotalBounties     int      `json:"total_bounties"`
	AvailableBounties int      `json:"available_bounties"`
	Bounties          []Bounty `json:"bounties"`
}

type SetupRequest struct {
	InstallationID int64  `json:"installation_id"`
	WalletAddress  string `json:"wallet_address"`
}

package handlers

type BountyResponse struct {
	Bounties []Bounty `json:"bounties"`
}

type Bounty struct {
	OwnerID        int64  `json:"owner_id"`
	Owner          string `json:"owner"`
	OwnerURL       string `json:"owner_url"`
	OwnerAvatarURL string `json:"owner_avatar_url"`
	OwnerType      string `json:"owner_type"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Reward         string `json:"reward"`
}

type SetupRequest struct {
	InstallationID int64  `json:"installation_id"`
	WalletAddress  string `json:"wallet_address"`
}

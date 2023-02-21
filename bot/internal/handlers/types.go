package handlers

type BountyResponse struct {
	Bounties []Bounty `json:"bounties"`
}

type Bounty struct {
	OwnerID        int64  `json:"owner_id"`
	Owner          string `json:"owner"`
	OwnerAvatarURL string `json:"owner_avatar_url"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Reward         uint64 `json:"reward"`
}

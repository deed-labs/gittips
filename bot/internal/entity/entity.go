package entity

type Owner struct {
	ID              int64
	Login           string
	URL             string
	AvatarURL       string
	TwitterUsername string
}

type Bounty struct {
	OwnerID       int64
	Title         string
	URL           string
	WalletAddress string
	Reward        uint64

	Owner Owner
}

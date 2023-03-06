package entity

import "math/big"

type Owner struct {
	ID              int64
	Login           string
	URL             string
	AvatarURL       string
	Type            string
	TwitterUsername string
	WalletAddress   string
}

type Bounty struct {
	ID             int64
	OwnerID        int64
	OwnerLogin     string
	OwnerURL       string
	OwnerAvatarURL string
	OwnerType      string
	Title          string
	URL            string
	Reward         *big.Int
}

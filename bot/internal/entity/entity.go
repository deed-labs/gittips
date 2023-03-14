package entity

import "math/big"

type Owner struct {
	ID              int64
	Name            string
	Login           string
	URL             string
	AvatarURL       string
	Type            string
	TwitterUsername string
	WalletAddress   string
}

type BountyWithOwner struct {
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

type Bounty struct {
	ID      int64
	OwnerID int64
	Title   string
	URL     string
	Reward  *big.Int
	Closed  bool
}

type InstallationInfo struct {
	Installed bool
	OwnerName string
	OwnerID   int64
}

type OwnerFullInfo struct {
	Owner             *Owner
	TotalBudget       *big.Int
	AvailableBudget   *big.Int
	TotalBounties     int
	AvailableBounties int
	Bounties          []*Bounty
}

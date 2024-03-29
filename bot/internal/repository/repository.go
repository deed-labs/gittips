package repository

import (
	"context"
	"errors"
	"math/big"

	"github.com/deed-labs/gittips/bot/internal/entity"
)

type Repository interface {
	Owners() OwnersRepository
	Bounties() BountiesRepository
}

type OwnersRepository interface {
	Get(ctx context.Context, ownerID int64) (*entity.Owner, error)
	GetByWalletAddress(ctx context.Context, address string) (*entity.Owner, error)
	Save(ctx context.Context, owner *entity.Owner) error

	SetWalletAddress(ctx context.Context, ownerId int64, walletAddress string) error
}

type BountiesRepository interface {
	GetAll(ctx context.Context) ([]*entity.BountyWithOwner, error)
	GetByOwnerId(ctx context.Context, ownerId int64) ([]*entity.Bounty, error)
	Save(ctx context.Context, bounty *entity.Bounty) error
	Delete(ctx context.Context, id int64) error

	SetReward(ctx context.Context, bountyId int64, value *big.Int) error
	SetClosed(ctx context.Context, bountyId int64, closed bool) error
}

var ErrNotFound = errors.New("not found")

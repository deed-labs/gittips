package repository

import (
	"context"
	"errors"
	"github.com/deed-labs/openroll/bot/internal/entity"
)

type Repository interface {
	Owners() OwnersRepository
	Bounties() BountiesRepository
}

type OwnersRepository interface {
	Get(ctx context.Context, ownerID int64) (*entity.Owner, error)
	Save(ctx context.Context, owner *entity.Owner) error
}

type BountiesRepository interface {
	GetAll(ctx context.Context) ([]*entity.Bounty, error)
	Save(ctx context.Context, bounty *entity.Bounty) error
}

var ErrNotFound = errors.New("not found")

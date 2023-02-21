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
}

type BountiesRepository interface {
	Create(ctx context.Context, bounty *entity.Bounty) error
}

var ErrNotFound = errors.New("not found")

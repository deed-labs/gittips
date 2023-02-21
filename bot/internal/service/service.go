package service

import (
	"context"
	"github.com/deed-labs/openroll/bot/internal/entity"
	"github.com/deed-labs/openroll/bot/internal/repository"
)

type Owners interface {
	Exists(ctx context.Context, ID int64) (bool, error)
	Create(ctx context.Context, ID int64, login string, url string) error
}

type Bounties interface {
	GetAll(ctx context.Context) ([]*entity.Bounty, error)
	Create(ctx context.Context, ownerID int64, title string, url string, body string) error
}

type Chain interface {
}

type Services struct {
	Owners   Owners
	Bounties Bounties
}

type Deps struct {
	Repository repository.Repository
}

func New(deps *Deps) *Services {
	ownersSvc := NewOwnersService(deps.Repository)
	bountiesSvc := NewBountiesService(ownersSvc, deps.Repository)

	return &Services{
		Owners:   ownersSvc,
		Bounties: bountiesSvc,
	}
}

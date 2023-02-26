package service

import (
	"context"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
)

type Owners interface {
	Exists(ctx context.Context, id int64) (bool, error)
	Create(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error
}

type Bounties interface {
	GetAll(ctx context.Context) ([]*entity.Bounty, error)
	Create(ctx context.Context, id int64, ownerID int64, title string, url string, body string) error
	Delete(ctx context.Context, id int64) error
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

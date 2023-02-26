package service

import (
	"context"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/parser"
	"github.com/deed-labs/gittips/bot/internal/repository"
)

type BountiesService struct {
	owners Owners

	repository repository.Repository
}

func NewBountiesService(owners Owners, repository repository.Repository) *BountiesService {
	return &BountiesService{
		owners:     owners,
		repository: repository,
	}
}

func (s *BountiesService) GetAll(ctx context.Context) ([]*entity.Bounty, error) {
	return s.repository.Bounties().GetAll(ctx)
}

func (s *BountiesService) Create(ctx context.Context, id int64, ownerID int64, title string, url string, body string) error {
	parsedBody := parser.ParseBody(body)

	bounty := &entity.Bounty{
		ID:      id,
		OwnerID: ownerID,
		Title:   title,
		URL:     url,
		Reward:  parsedBody.Reward,
	}

	return s.repository.Bounties().Save(ctx, bounty)
}

func (s *BountiesService) Delete(ctx context.Context, id int64) error {
	return s.repository.Bounties().Delete(ctx, id)
}

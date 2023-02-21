package service

import (
	"context"
	"github.com/deed-labs/openroll/bot/internal/entity"
	"github.com/deed-labs/openroll/bot/internal/parser"
	"github.com/deed-labs/openroll/bot/internal/repository"
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

func (s *BountiesService) Create(ctx context.Context, ownerID int64, title string, url string, body string) error {
	parsedBody := parser.ParseBody(body)

	bounty := &entity.Bounty{
		OwnerID:       ownerID,
		Title:         title,
		URL:           url,
		WalletAddress: parsedBody.WalletAddress,
		Reward:        parsedBody.Reward,
	}

	return s.repository.Bounties().Create(ctx, bounty)
}

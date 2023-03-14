package service

import (
	"context"
	"strings"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/parser"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/xssnick/tonutils-go/tlb"
)

type BountiesService struct {
	repository repository.Repository
}

func NewBountiesService(repository repository.Repository) *BountiesService {
	return &BountiesService{
		repository: repository,
	}
}

func (s *BountiesService) GetAll(ctx context.Context) ([]*entity.BountyWithOwner, error) {
	return s.repository.Bounties().GetAll(ctx)
}

func (s *BountiesService) GetByOwnerId(ctx context.Context, ownerId int64) ([]*entity.Bounty, error) {
	return s.repository.Bounties().GetByOwnerId(ctx, ownerId)
}

func (s *BountiesService) Create(ctx context.Context, id int64, ownerID int64, title string, url string, body string) error {
	parsedBody := parser.Parse(body)

	var reward string
	if parsedBody.Reward != "" {
		reward = strings.Replace(parsedBody.Reward, ",", ".", 1)
	} else {
		reward = "0"
	}

	parsedReward, err := tlb.FromTON(reward)
	if err != nil {
		return ErrInvalidValue
	}

	// TODO: check if budget balance is sufficient for this reward

	bounty := &entity.Bounty{
		ID:      id,
		OwnerID: ownerID,
		Title:   title,
		URL:     url,
		Reward:  parsedReward.NanoTON(),
		Closed:  false,
	}

	return s.repository.Bounties().Save(ctx, bounty)
}

func (s *BountiesService) Delete(ctx context.Context, id int64) error {
	return s.repository.Bounties().Delete(ctx, id)
}

func (s *BountiesService) Close(ctx context.Context, id int64) error {
	return s.repository.Bounties().SetClosed(ctx, id, true)
}

package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/parser"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/xssnick/tonutils-go/tlb"
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
	parsedBody := parser.Parse(body)

	parsedBody.Reward = strings.Replace(parsedBody.Reward, ",", ".", 1)

	parsedReward, err := tlb.FromTON(parsedBody.Reward)
	if err != nil {
		return fmt.Errorf("parse reward amount: %w", err)
	}

	bounty := &entity.Bounty{
		ID:      id,
		OwnerID: ownerID,
		Title:   title,
		URL:     url,
		Reward:  parsedReward.NanoTON(),
	}

	return s.repository.Bounties().Save(ctx, bounty)
}

func (s *BountiesService) Delete(ctx context.Context, id int64) error {
	return s.repository.Bounties().Delete(ctx, id)
}

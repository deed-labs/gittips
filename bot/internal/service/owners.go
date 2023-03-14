package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/deed-labs/gittips/bot/internal/ton"
)

type OwnersService struct {
	Bounties Bounties
	ton      *ton.TON

	repository repository.Repository
}

func NewOwnersService(bounties Bounties, ton *ton.TON, repository repository.Repository) *OwnersService {
	return &OwnersService{
		Bounties:   bounties,
		ton:        ton,
		repository: repository,
	}
}

func (s *OwnersService) Exists(ctx context.Context, id int64) (bool, error) {
	_, err := s.repository.Owners().Get(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("get owner: %w", err)
	}

	return true, nil
}

func (s *OwnersService) Get(ctx context.Context, id int64) (*entity.OwnerFullInfo, error) {
	owner, err := s.repository.Owners().Get(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrOwnerNotFound
		}

		return nil, fmt.Errorf("get owner: %w", err)
	}

	balance, err := s.ton.GetBudgetBalance(ctx, owner.WalletAddress)
	if err != nil {
		return nil, fmt.Errorf("get budget balance: %w", err)
	}

	allBounties, err := s.Bounties.GetByOwnerId(ctx, owner.ID)
	if err != nil {
		return nil, fmt.Errorf("get bounties: %w", err)
	}
	availableBounties := make([]*entity.Bounty, 0, len(allBounties))

	availableBudget := new(big.Int).Set(balance)
	for _, bounty := range allBounties {
		if bounty.Closed {
			continue
		}

		availableBudget.Sub(availableBudget, bounty.Reward)
		availableBounties = append(availableBounties, bounty)
	}

	info := &entity.OwnerFullInfo{
		Owner:             owner,
		TotalBudget:       balance,
		AvailableBudget:   availableBudget,
		TotalBounties:     len(allBounties),
		AvailableBounties: len(availableBounties),
		Bounties:          availableBounties,
	}

	return info, nil
}

func (s *OwnersService) Create(ctx context.Context, id int64, login string, name string, url string, avatarURL string,
	ownerType string) error {
	owner := &entity.Owner{
		ID:              id,
		Name:            name,
		Login:           login,
		URL:             url,
		AvatarURL:       avatarURL,
		Type:            ownerType,
		TwitterUsername: "",
		WalletAddress:   "",
	}

	return s.repository.Owners().Save(ctx, owner)
}

func (s *OwnersService) LinkWithWallet(ctx context.Context, ownerId int64, walletAddress string) error {
	return s.repository.Owners().SetWalletAddress(ctx, ownerId, walletAddress)
}

func (s *OwnersService) GetInstallationInfo(ctx context.Context, address string) (*entity.InstallationInfo, error) {
	owner, err := s.repository.Owners().GetByWalletAddress(ctx, address)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return &entity.InstallationInfo{Installed: false}, nil
		}

		return nil, fmt.Errorf("get owner: %w", err)
	}

	info := &entity.InstallationInfo{
		Installed: true,
		OwnerName: owner.Name,
		OwnerID:   owner.ID,
	}

	return info, nil
}

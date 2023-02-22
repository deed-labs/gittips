package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/deed-labs/openroll/bot/internal/entity"
	"github.com/deed-labs/openroll/bot/internal/repository"
)

type OwnersService struct {
	repository repository.Repository
}

func NewOwnersService(repository repository.Repository) *OwnersService {
	return &OwnersService{repository: repository}
}

func (s *OwnersService) Exists(ctx context.Context, ID int64) (bool, error) {
	_, err := s.repository.Owners().Get(ctx, ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("get owner: %w", err)
	}

	return true, nil
}

func (s *OwnersService) Create(ctx context.Context, ID int64, login string, url string, avatarURL string) error {
	owner := &entity.Owner{
		ID:              ID,
		Login:           login,
		URL:             url,
		AvatarURL:       avatarURL,
		TwitterUsername: "",
	}

	return s.repository.Owners().Save(ctx, owner)
}

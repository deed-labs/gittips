package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
)

type OwnersService struct {
	repository repository.Repository
}

func NewOwnersService(repository repository.Repository) *OwnersService {
	return &OwnersService{repository: repository}
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

func (s *OwnersService) Create(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error {
	owner := &entity.Owner{
		ID:              id,
		Login:           login,
		URL:             url,
		AvatarURL:       avatarURL,
		Type:            ownerType,
		TwitterUsername: "",
	}

	return s.repository.Owners().Save(ctx, owner)
}

package service

import (
	"context"
	"github.com/deed-labs/openroll/bot/internal/repository"
)

type OwnersService struct {
	repository repository.Repository
}

func NewOwnersService(repository repository.Repository) *OwnersService {
	return &OwnersService{repository: repository}
}

func (s *OwnersService) Exists(ctx context.Context, ID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OwnersService) Create(ctx context.Context, ID int64, login string, url string) error {
	//TODO implement me
	panic("implement me")
}

package service

import (
	"context"
	"github.com/deed-labs/gittips/bot/internal/ton"
)

type CommentsService struct {
	ton *ton.TON
}

func NewCommentsService(ton *ton.TON) *CommentsService {
	return &CommentsService{ton: ton}
}

func (c *CommentsService) Process(ctx context.Context, senderId int64, ownerId int64, body string) error {
	//TODO implement me
	panic("implement me")
}

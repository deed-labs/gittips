package service

import (
	"context"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/deed-labs/gittips/bot/internal/ton"
	"strings"
)

type CommandsService struct {
	ton        *ton.TON
	repository repository.Repository
}

func NewCommandsService(ton *ton.TON, repository repository.Repository) *CommandsService {
	return &CommandsService{
		ton:        ton,
		repository: repository,
	}
}

func (s *CommandsService) Process(ctx context.Context, ownerId int64, commands []string) error {
	for _, v := range commands {
		cmd := strings.Fields(v)
		if len(cmd) == 0 {
			// invalid command
			continue
		}

		var err error
		switch cmd[0] {
		case "pay":
			err = s.sendPayout(ctx, ownerId)
		default:
			// unknown command
			continue
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CommandsService) sendPayout(ctx context.Context, ownerId int64) error {
	// TODO
	return nil
}

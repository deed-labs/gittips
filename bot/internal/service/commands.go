package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/deed-labs/gittips/bot/internal/ton"
	"github.com/xssnick/tonutils-go/tlb"
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

func (s *CommandsService) Parse(command string) interface{} {
	cmd := strings.Fields(command)
	if len(cmd) == 0 {
		// invalid command
		return nil
	}

	switch cmd[0] {
	case "pay":
		return &SendPaymentCommand{
			svc:   s,
			To:    strings.TrimPrefix(cmd[1], "@"),
			Value: cmd[2],
		}
	case "set":
		if len(cmd) < 3 {
			return nil
		}

		switch strings.ToLower(cmd[1]) {
		case "wallet":
			return &SetWalletCommand{
				svc:           s,
				WalletAddress: cmd[2],
			}
		case "reward":
			return &SetRewardCommand{
				svc:         s,
				RewardValue: cmd[2],
			}
		default:
			return nil
		}
	case "close":
		return &CloseCommand{}
	default:
		// unknown command
		return nil
	}
}

type SendPaymentCommand struct {
	svc   *CommandsService
	To    string
	Value string
}

func (c *SendPaymentCommand) Run(ctx context.Context, fromOwnerId int64, toOwnerId int64, value string) error {
	fromOwner, err := c.svc.repository.Owners().Get(ctx, fromOwnerId)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return ErrOwnerNotFound
	} else if err != nil {
		return fmt.Errorf("get owner: %w", err)
	}

	toOwner, err := c.svc.repository.Owners().Get(ctx, toOwnerId)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return ErrUserNotFound
	} else if err != nil {
		return fmt.Errorf("get owner: %w", err)
	}

	tonValue, err := tlb.FromTON(strings.Replace(value, ",", ".", 1))
	if err != nil {
		return ErrInvalidValue
	}

	if err := c.svc.ton.SendPayout(ctx, fromOwner.WalletAddress, toOwner.WalletAddress, tonValue.NanoTON()); err != nil {
		return fmt.Errorf("send payout: %w", err)
	}

	return nil
}

type SetWalletCommand struct {
	svc           *CommandsService
	WalletAddress string
}

func (c *SetWalletCommand) Run(ctx context.Context, ownerId int64, address string) error {
	return c.svc.repository.Owners().SetWalletAddress(ctx, ownerId, address)
}

type SetRewardCommand struct {
	svc         *CommandsService
	RewardValue string
}

func (c *SetRewardCommand) Run(ctx context.Context, bountyId int64, newReward string) error {
	tonValue, err := tlb.FromTON(strings.Replace(newReward, ",", ".", 1))
	if err != nil {
		return ErrInvalidValue
	}

	return c.svc.repository.Bounties().SetReward(ctx, bountyId, tonValue.NanoTON())
}

type CloseCommand struct{}

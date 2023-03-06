package ton

import (
	"context"
	"fmt"
	"math/big"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type TON struct {
	client     *ton.APIClient
	wallet     *wallet.Wallet
	routerAddr *address.Address
}

const sendPayoutOP = 3

func New(client *ton.APIClient, wallet *wallet.Wallet, routerAddr string) *TON {
	return &TON{
		client:     client,
		wallet:     wallet,
		routerAddr: address.MustParseAddr(routerAddr),
	}
}

func (t *TON) SendPayout(ctx context.Context, to string, amount *big.Int) error {
	destination, err := address.ParseAddr(to)
	if err != nil {
		return fmt.Errorf("parse destination: %w", err)
	}

	body := cell.BeginCell().
		MustStoreUInt(sendPayoutOP, 32). // op code
		MustStoreUInt(0, 64).            // query id
		MustStoreAddr(destination).
		MustStoreCoins(amount.Uint64()).
		EndCell()

	err = t.wallet.Send(ctx, &wallet.Message{
		Mode: 1, // pay fees separately (from balance, not from amount)
		InternalMessage: &tlb.InternalMessage{
			DstAddr: t.routerAddr,
			Amount:  tlb.MustFromTON("0.03"),
			Body:    body,
		},
	}, true)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

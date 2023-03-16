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

var zeroBigInt = new(big.Int).SetInt64(0)

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

func (t *TON) GetBudgetBalance(ctx context.Context, walletAddress string) (*big.Int, error) {
	walletAddr, err := address.ParseAddr(walletAddress)
	if err != nil {
		return nil, fmt.Errorf("parse address: %w", err)
	}

	ctx = t.client.Client().StickyContext(ctx)

	block, err := t.client.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("get masterchain info: %w", err)
	}

	param := cell.BeginCell().MustStoreAddr(walletAddr).EndCell()

	res, err := t.client.RunGetMethod(ctx, block, t.routerAddr, "get_budget_address", param.BeginParse())
	if err != nil {
		return nil, fmt.Errorf("run get method: %w", err)
	}

	sc, err := res.Slice(0)
	if err != nil {
		return nil, fmt.Errorf("load slice from result: %w", err)
	}
	budgetAddr := sc.MustLoadAddr()

	account, err := t.client.GetAccount(ctx, block, budgetAddr)
	if err != nil {
		return nil, fmt.Errorf("get account: %w", err)
	}

	if !account.IsActive {
		return zeroBigInt, nil
	}

	return account.State.Balance.NanoTON(), nil
}

package ton

import (
	"context"
	"errors"
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

func (t *TON) SendPayout(ctx context.Context, ownerAddr string, toAddr string, amount *big.Int) error {
	owner, err := address.ParseAddr(ownerAddr)
	if err != nil {
		return fmt.Errorf("parse owner address: %w", err)
	}

	destination, err := address.ParseAddr(toAddr)
	if err != nil {
		return fmt.Errorf("parse destination address: %w", err)
	}

	body := cell.BeginCell().
		MustStoreUInt(sendPayoutOP, 32). // op code
		MustStoreUInt(0, 64).            // query id
		MustStoreAddr(owner).
		MustStoreAddr(destination).
		MustStoreCoins(amount.Uint64()).
		EndCell()

	err = t.wallet.Send(ctx, &wallet.Message{
		Mode: 1, // pay fees separately (from balance, not from amount)
		InternalMessage: &tlb.InternalMessage{
			Bounce:  true,
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
	var execError ton.ContractExecError
	if err != nil && errors.As(err, &execError) && execError.Code == ton.ErrCodeContractNotInitialized {
		// Returns a zero balance if the error is only caused by
		// the fact that the budget contract has not yet been initialized.
		return zeroBigInt, nil
	}
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

	return account.State.Balance.NanoTON(), nil
}

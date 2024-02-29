package db

import (
	"context"
	"testing"
	"time"

	"github.com/kahakai/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, account1, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotZero(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestListTransfers(t *testing.T) {
	var accounts1 []Account
	var accounts2 []Account

	for i := 0; i < 10; i++ {
		accounts1 = append(accounts1, createRandomAccount(t))
	}

	for i := 0; i < 10; i++ {
		accounts2 = append(accounts2, createRandomAccount(t))
	}

	for _, account1 := range accounts1 {
		for _, account2 := range accounts2 {
			for i := 0; i < 10; i++ {
				createRandomTransfer(t, account1, account2)
			}

			for i := 0; i < 10; i++ {
				createRandomTransfer(t, account2, account1)
			}
		}
	}

	for _, account1 := range accounts1 {
		for _, account2 := range accounts2 {
			arg1 := ListTransfersParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Limit:         5,
				Offset:        5,
			}

			transfers1, err := testQueries.ListTransfers(context.Background(), arg1)
			require.NoError(t, err)
			require.Len(t, transfers1, 5)

			for _, transfer := range transfers1 {
				require.NotEmpty(t, transfer)
			}

			arg2 := ListTransfersParams{
				FromAccountID: account2.ID,
				ToAccountID:   account1.ID,
				Limit:         5,
				Offset:        5,
			}

			transfers2, err := testQueries.ListTransfers(context.Background(), arg2)
			require.NoError(t, err)
			require.Len(t, transfers2, 5)

			for _, transfer := range transfers2 {
				require.NotEmpty(t, transfer)
			}
		}
	}
}

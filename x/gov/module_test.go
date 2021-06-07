package gov_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abcitypes "github.com/mydexchain/tendermint/abci/types"
	tmproto "github.com/mydexchain/tendermint/proto/tendermint/types"

	"github.com/mydexchain/chain-sdk/simapp"
	authtypes "github.com/mydexchain/chain-sdk/x/auth/types"
	"github.com/mydexchain/chain-sdk/x/gov/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}

package gov_test

import (
	"strings"
	"testing"

	tmproto "github.com/mydexchain/tendermint/proto/tendermint/types"

	"github.com/mydexchain/chain-sdk/testutil/testdata"

	"github.com/stretchr/testify/require"

	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/x/gov"
	"github.com/mydexchain/chain-sdk/x/gov/keeper"
)

func TestInvalidMsg(t *testing.T) {
	k := keeper.Keeper{}
	h := gov.NewHandler(k)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized gov message type"))
}

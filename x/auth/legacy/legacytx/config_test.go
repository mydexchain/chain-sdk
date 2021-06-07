package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mydexchain/chain-sdk/codec"
	cryptoAmino "github.com/mydexchain/chain-sdk/crypto/codec"
	"github.com/mydexchain/chain-sdk/testutil/testdata"
	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/x/auth/legacy/legacytx"
	"github.com/mydexchain/chain-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "chain-sdk/TEst", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}

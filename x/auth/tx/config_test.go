package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mydexchain/chain-sdk/codec"
	codectypes "github.com/mydexchain/chain-sdk/codec/types"
	"github.com/mydexchain/chain-sdk/std"
	"github.com/mydexchain/chain-sdk/testutil/testdata"
	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}

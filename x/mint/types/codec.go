package types

import (
	"github.com/mydexchain/chain-sdk/codec"
	cryptocodec "github.com/mydexchain/chain-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

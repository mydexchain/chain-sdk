package legacytx

import (
	"github.com/mydexchain/chain-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "chain-sdk/STdtX", nil)
}

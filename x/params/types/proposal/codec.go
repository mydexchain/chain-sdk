package proposal

import (
	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/codec/types"
	govtypes "github.com/mydexchain/chain-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers all necessary param module types with a given LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&ParameterChangeProposal{}, "chain-sdk/PArametercHangepRoposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ParameterChangeProposal{},
	)
}

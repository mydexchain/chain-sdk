package types

import (
	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/codec/types"
	govtypes "github.com/mydexchain/chain-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(Plan{}, "chain-sdk/PLan", nil)
	cdc.RegisterConcrete(&SoftwareUpgradeProposal{}, "chain-sdk/SOftwareuPgradepRoposal", nil)
	cdc.RegisterConcrete(&CancelSoftwareUpgradeProposal{}, "chain-sdk/CAncelsOftwareuPgradepRoposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SoftwareUpgradeProposal{},
		&CancelSoftwareUpgradeProposal{},
	)
}

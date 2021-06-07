package types

import (
	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/codec/types"
	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/types/msgservice"
	authtypes "github.com/mydexchain/chain-sdk/x/auth/types"
	"github.com/mydexchain/chain-sdk/x/auth/vesting/exported"
)

// RegisterLegacyAminoCodec registers the vesting interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseVestingAccount{}, "chain-sdk/BAsevEstingaCcount", nil)
	cdc.RegisterConcrete(&ContinuousVestingAccount{}, "chain-sdk/COntinuousvEstingaCcount", nil)
	cdc.RegisterConcrete(&DelayedVestingAccount{}, "chain-sdk/DElayedvEstingaCcount", nil)
	cdc.RegisterConcrete(&PeriodicVestingAccount{}, "chain-sdk/PEriodicvEstingaCcount", nil)
}

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"dex.vesting.v0.VestingAccount",
		(*exported.VestingAccount)(nil),
		&ContinuousVestingAccount{},
		&DelayedVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

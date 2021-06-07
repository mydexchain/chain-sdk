package types

import (
	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/codec/types"
	cryptocodec "github.com/mydexchain/chain-sdk/crypto/codec"
	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/types/msgservice"
	"github.com/mydexchain/chain-sdk/x/bank/exported"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.SupplyI)(nil), nil)
	cdc.RegisterConcrete(&Supply{}, "chain-sdk/SUpply", nil)
	cdc.RegisterConcrete(&MsgSend{}, "chain-sdk/MSgsEnd", nil)
	cdc.RegisterConcrete(&MsgMultiSend{}, "chain-sdk/MSgmUltisEnd", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)

	registry.RegisterInterface(
		"dex.bank.v0.SupplyI",
		(*exported.SupplyI)(nil),
		&Supply{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/bank module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

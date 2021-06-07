package types

import (
	codectypes "github.com/mydexchain/chain-sdk/codec/types"
	clienttypes "github.com/mydexchain/chain-sdk/x/ibc/core/02-client/types"
	connectiontypes "github.com/mydexchain/chain-sdk/x/ibc/core/03-connection/types"
	channeltypes "github.com/mydexchain/chain-sdk/x/ibc/core/04-channel/types"
	commitmenttypes "github.com/mydexchain/chain-sdk/x/ibc/core/23-commitment/types"
	solomachinetypes "github.com/mydexchain/chain-sdk/x/ibc/light-clients/06-solomachine/types"
	ibctmtypes "github.com/mydexchain/chain-sdk/x/ibc/light-clients/07-tendermint/types"
	localhosttypes "github.com/mydexchain/chain-sdk/x/ibc/light-clients/09-localhost/types"
)

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	ibctmtypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}

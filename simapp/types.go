package simapp

import (
	abci "github.com/mydexchain/tendermint/abci/types"

	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/server/types"
	sdk "github.com/mydexchain/chain-sdk/types"
	"github.com/mydexchain/chain-sdk/types/module"
)

// App implements the common methods for a Chain SDK-based application
// specific blockchain.
type App interface {
	// The assigned name of the app.
	Name() string

	// The application types codec.
	// NOTE: This shoult be sealed before being returned.
	LegacyAmino() *codec.LegacyAmino

	// Application updates every begin block.
	BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock

	// Application updates every end block.
	EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock

	// Application update at chain (i.e app) initialization.
	InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain

	// Loads the app at a given height.
	LoadHeight(height int64) error

	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs []string,
	) (types.ExportedApp, error)

	// All the registered module account addreses.
	ModuleAccountAddrs() map[string]bool

	// Helper for the simulation framework.
	SimulationManager() *module.SimulationManager
}

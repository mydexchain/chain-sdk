// DONTCOVER
// nolint
package v036

import (
	sdk "github.com/mydexchain/chain-sdk/types"
)

const ModuleName = "supply"

type (
	GenesisState struct {
		Supply sdk.Coins `json:"supply" yaml:"supply"`
	}
)

func EmptyGenesisState() GenesisState {
	return GenesisState{
		Supply: sdk.NewCoins(), // leave this empty as it's filled on initialization
	}
}

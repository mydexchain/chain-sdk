// DONTCOVER
// nolint
package v036

import v034auth "github.com/mydexchain/chain-sdk/x/auth/legacy/v034"

const (
	ModuleName = "auth"
)

type (
	GenesisState struct {
		Params v034auth.Params `json:"params"`
	}
)

func NewGenesisState(params v034auth.Params) GenesisState {
	return GenesisState{params}
}
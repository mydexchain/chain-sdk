package cli

import (
	"io/ioutil"

	"github.com/mydexchain/chain-sdk/codec"
	"github.com/mydexchain/chain-sdk/x/distribution/types"
)

// ParseCommunityPoolSpendProposalWithDeposit reads and parses a CommunityPoolSpendProposalWithDeposit from a file.
func ParseCommunityPoolSpendProposalWithDeposit(cdc codec.JSONMarshaler, proposalFile string) (types.CommunityPoolSpendProposalWithDeposit, error) {
	proposal := types.CommunityPoolSpendProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
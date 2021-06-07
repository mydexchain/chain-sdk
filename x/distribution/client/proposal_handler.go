package client

import (
	"github.com/mydexchain/chain-sdk/x/distribution/client/cli"
	"github.com/mydexchain/chain-sdk/x/distribution/client/rest"
	govclient "github.com/mydexchain/chain-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)

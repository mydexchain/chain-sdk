package client

import (
	govclient "github.com/mydexchain/chain-sdk/x/gov/client"
	"github.com/mydexchain/chain-sdk/x/params/client/cli"
	"github.com/mydexchain/chain-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)

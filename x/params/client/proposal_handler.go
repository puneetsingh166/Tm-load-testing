package client

import (
	govclient "github.com/puneetsingh166/tm-load-test/x/gov/client"
	"github.com/puneetsingh166/tm-load-test/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)

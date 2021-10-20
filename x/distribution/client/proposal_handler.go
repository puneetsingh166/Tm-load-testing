package client

import (
	"github.com/puneetsingh166/tm-load-test/x/distribution/client/cli"
	govclient "github.com/puneetsingh166/tm-load-test/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)

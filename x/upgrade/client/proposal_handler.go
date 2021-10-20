package client

import (
	govclient "github.com/puneetsingh166/tm-load-test/x/gov/client"
	"github.com/puneetsingh166/tm-load-test/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)

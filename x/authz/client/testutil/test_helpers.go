package testutil

import (
	"github.com/puneetsingh166/tm-load-test/testutil"
	clitestutil "github.com/puneetsingh166/tm-load-test/testutil/cli"
	"github.com/puneetsingh166/tm-load-test/testutil/network"
	"github.com/puneetsingh166/tm-load-test/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}

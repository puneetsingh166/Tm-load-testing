package tmservice

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/puneetsingh166/tm-load-test/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}

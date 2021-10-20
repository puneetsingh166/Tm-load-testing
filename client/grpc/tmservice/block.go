package tmservice

import (
	"context"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/puneetsingh166/tm-load-test/client"
)

func getBlock(ctx context.Context, clientCtx client.Context, height *int64) (*ctypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(ctx, height)
}

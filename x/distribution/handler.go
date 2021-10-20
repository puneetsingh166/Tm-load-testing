package distribution

import (
	sdk "github.com/puneetsingh166/tm-load-test/types"
	sdkerrors "github.com/puneetsingh166/tm-load-test/types/errors"
	"github.com/puneetsingh166/tm-load-test/x/distribution/keeper"
	"github.com/puneetsingh166/tm-load-test/x/distribution/types"
	govtypes "github.com/puneetsingh166/tm-load-test/x/gov/types"
)

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}

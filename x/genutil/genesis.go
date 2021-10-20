package genutil

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/puneetsingh166/tm-load-test/client"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []abci.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
	}
	return
}

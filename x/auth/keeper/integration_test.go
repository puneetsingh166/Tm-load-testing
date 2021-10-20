package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/puneetsingh166/tm-load-test/simapp"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	authtypes "github.com/puneetsingh166/tm-load-test/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(t *testing.T, isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}

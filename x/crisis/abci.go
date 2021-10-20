package crisis

import (
	"time"

	"github.com/puneetsingh166/tm-load-test/telemetry"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/x/crisis/keeper"
	"github.com/puneetsingh166/tm-load-test/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)
}

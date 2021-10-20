package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/puneetsingh166/tm-load-test/simapp"
	"github.com/puneetsingh166/tm-load-test/types/module"
	simtypes "github.com/puneetsingh166/tm-load-test/types/simulation"
	"github.com/puneetsingh166/tm-load-test/x/feegrant"
	"github.com/puneetsingh166/tm-load-test/x/feegrant/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	app := simapp.Setup(t, false)

	s := rand.NewSource(1)
	r := rand.New(s)

	accounts := simtypes.RandomAccounts(r, 3)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          app.AppCodec(),
		Rand:         r,
		NumBonded:    3,
		Accounts:     accounts,
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var feegrantGenesis feegrant.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[feegrant.ModuleName], &feegrantGenesis)

	require.Len(t, feegrantGenesis.Allowances, len(accounts)-1)
}

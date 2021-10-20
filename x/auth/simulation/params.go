package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	simtypes "github.com/puneetsingh166/tm-load-test/types/simulation"
	"github.com/puneetsingh166/tm-load-test/x/auth/types"
	"github.com/puneetsingh166/tm-load-test/x/simulation"
)

const (
	keyMaxMemoCharacters = "MaxMemoCharacters"
	keyTxSigLimit        = "TxSigLimit"
	keyTxSizeCostPerByte = "TxSizeCostPerByte"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyMaxMemoCharacters,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenMaxMemoChars(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyTxSigLimit,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenTxSigLimit(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyTxSizeCostPerByte,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenTxSizeCostPerByte(r))
			},
		),
	}
}

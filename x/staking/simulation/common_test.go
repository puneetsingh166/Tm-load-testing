package simulation_test

import (
	"math/big"

	sdk "github.com/puneetsingh166/tm-load-test/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

package keeper_test

import sdk "github.com/puneetsingh166/tm-load-test/types"

var (
	// The default power validators are initialized to have within tests
	InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
)

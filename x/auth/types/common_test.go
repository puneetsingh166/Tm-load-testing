package types_test

import (
	"github.com/puneetsingh166/tm-load-test/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)

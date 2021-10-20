package legacytx

import (
	"github.com/puneetsingh166/tm-load-test/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}

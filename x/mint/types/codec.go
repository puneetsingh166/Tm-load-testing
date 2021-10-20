package types

import (
	"github.com/puneetsingh166/tm-load-test/codec"
	cryptocodec "github.com/puneetsingh166/tm-load-test/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

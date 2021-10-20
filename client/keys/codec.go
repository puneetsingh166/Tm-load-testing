package keys

import (
	"github.com/puneetsingh166/tm-load-test/codec"
	cryptocodec "github.com/puneetsingh166/tm-load-test/crypto/codec"
)

// TODO: remove this file https://github.com/puneetsingh166/tm-load-test/issues/8047

// KeysCdc defines codec to be used with key operations
var KeysCdc *codec.LegacyAmino

func init() {
	KeysCdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(KeysCdc)
	KeysCdc.Seal()
}

// marshal keys
func MarshalJSON(o interface{}) ([]byte, error) {
	return KeysCdc.MarshalJSON(o)
}

// unmarshal json
func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return KeysCdc.UnmarshalJSON(bz, ptr)
}

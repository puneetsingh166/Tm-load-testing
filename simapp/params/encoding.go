package params

import (
	"github.com/puneetsingh166/tm-load-test/client"
	"github.com/puneetsingh166/tm-load-test/codec"
	"github.com/puneetsingh166/tm-load-test/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

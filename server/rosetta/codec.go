package rosetta

import (
	"github.com/puneetsingh166/tm-load-test/codec"
	codectypes "github.com/puneetsingh166/tm-load-test/codec/types"
	cryptocodec "github.com/puneetsingh166/tm-load-test/crypto/codec"
	authcodec "github.com/puneetsingh166/tm-load-test/x/auth/types"
	bankcodec "github.com/puneetsingh166/tm-load-test/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}

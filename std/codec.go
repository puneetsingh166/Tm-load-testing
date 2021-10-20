package std

import (
	"github.com/puneetsingh166/tm-load-test/codec"
	"github.com/puneetsingh166/tm-load-test/codec/types"
	cryptocodec "github.com/puneetsingh166/tm-load-test/crypto/codec"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	txtypes "github.com/puneetsingh166/tm-load-test/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}

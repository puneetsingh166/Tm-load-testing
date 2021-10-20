package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/puneetsingh166/tm-load-test/codec"
	codectypes "github.com/puneetsingh166/tm-load-test/codec/types"
	"github.com/puneetsingh166/tm-load-test/std"
	"github.com/puneetsingh166/tm-load-test/testutil/testdata"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}

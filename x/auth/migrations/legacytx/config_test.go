package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/puneetsingh166/tm-load-test/codec"
	cryptoAmino "github.com/puneetsingh166/tm-load-test/crypto/codec"
	"github.com/puneetsingh166/tm-load-test/testutil/testdata"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/x/auth/migrations/legacytx"
	"github.com/puneetsingh166/tm-load-test/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}

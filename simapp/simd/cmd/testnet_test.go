package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/puneetsingh166/tm-load-test/client"
	"github.com/puneetsingh166/tm-load-test/client/flags"
	"github.com/puneetsingh166/tm-load-test/server"
	"github.com/puneetsingh166/tm-load-test/simapp"
	banktypes "github.com/puneetsingh166/tm-load-test/x/bank/types"
	genutiltest "github.com/puneetsingh166/tm-load-test/x/genutil/client/testutil"
	genutiltypes "github.com/puneetsingh166/tm-load-test/x/genutil/types"
)

func Test_TestnetCmd(t *testing.T) {
	home := t.TempDir()
	encodingConfig := simapp.MakeTestEncodingConfig()
	logger := log.NewNopLogger()
	cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
	require.NoError(t, err)

	err = genutiltest.ExecInitCmd(simapp.ModuleBasics, home, encodingConfig.Codec)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithHomeDir(home).
		WithTxConfig(encodingConfig.TxConfig)

	ctx := context.Background()
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := testnetInitFilesCmd(simapp.ModuleBasics, banktypes.GenesisBalancesIterator{})
	cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend), fmt.Sprintf("--output-dir=%s", home)})
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)

	genFile := cfg.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)

	bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
	require.NotEmpty(t, bankGenState.Supply.String())
}

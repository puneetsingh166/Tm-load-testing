package testutil

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/suite"

	"github.com/puneetsingh166/tm-load-test/client"
	"github.com/puneetsingh166/tm-load-test/client/flags"
	"github.com/puneetsingh166/tm-load-test/simapp"
	"github.com/puneetsingh166/tm-load-test/testutil"
	"github.com/puneetsingh166/tm-load-test/testutil/network"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	banktypes "github.com/puneetsingh166/tm-load-test/x/bank/types"
	"github.com/puneetsingh166/tm-load-test/x/genutil/client/cli"
	stakingcli "github.com/puneetsingh166/tm-load-test/x/staking/client/cli"
	"github.com/puneetsingh166/tm-load-test/x/staking/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestGenTxCmd() {
	val := s.network.Validators[0]
	dir := s.T().TempDir()

	cmd := cli.GenTxCmd(
		simapp.ModuleBasics,
		val.ClientCtx.TxConfig, banktypes.GenesisBalancesIterator{}, val.ClientCtx.HomeDir)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	amount := sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(12))
	genTxFile := filepath.Join(dir, "myTx")
	cmd.SetArgs([]string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagOutputDocument, genTxFile),
		val.Moniker,
		amount.String(),
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	// validate generated transaction.
	open, err := os.Open(genTxFile)
	s.Require().NoError(err)

	all, err := ioutil.ReadAll(open)
	s.Require().NoError(err)

	tx, err := val.ClientCtx.TxConfig.TxJSONDecoder()(all)
	s.Require().NoError(err)

	msgs := tx.GetMsgs()
	s.Require().Len(msgs, 1)

	s.Require().Equal(sdk.MsgTypeURL(&types.MsgCreateValidator{}), sdk.MsgTypeURL(msgs[0]))
	s.Require().True(val.Address.Equals(msgs[0].GetSigners()[0]))
	s.Require().Equal(amount, msgs[0].(*types.MsgCreateValidator).Value)
	s.Require().NoError(tx.ValidateBasic())
}

func (s *IntegrationTestSuite) TestGenTxCmdPubkey() {
	val := s.network.Validators[0]
	dir := s.T().TempDir()

	cmd := cli.GenTxCmd(
		simapp.ModuleBasics,
		val.ClientCtx.TxConfig,
		banktypes.GenesisBalancesIterator{},
		val.ClientCtx.HomeDir,
	)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	amount := sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(12))
	genTxFile := filepath.Join(dir, "myTx")

	cmd.SetArgs([]string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagOutputDocument, genTxFile),
		fmt.Sprintf("--%s={\"key\":\"BOIkjkFruMpfOFC9oNPhiJGfmY2pHF/gwHdLDLnrnS0=\"}", stakingcli.FlagPubKey),
		val.Moniker,
		amount.String(),
	})
	s.Require().Error(cmd.ExecuteContext(ctx))

	cmd.SetArgs([]string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagOutputDocument, genTxFile),
		fmt.Sprintf("--%s={\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"BOIkjkFruMpfOFC9oNPhiJGfmY2pHF/gwHdLDLnrnS0=\"}", stakingcli.FlagPubKey),
		val.Moniker,
		amount.String(),
	})
	s.Require().NoError(cmd.ExecuteContext(ctx))
}

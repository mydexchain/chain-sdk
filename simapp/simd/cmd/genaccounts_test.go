package cmd_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/mydexchain/tendermint/libs/log"

	"github.com/mydexchain/chain-sdk/client"
	"github.com/mydexchain/chain-sdk/client/flags"
	"github.com/mydexchain/chain-sdk/server"
	"github.com/mydexchain/chain-sdk/simapp"
	simcmd "github.com/mydexchain/chain-sdk/simapp/simd/cmd"
	"github.com/mydexchain/chain-sdk/testutil/testdata"
	"github.com/mydexchain/chain-sdk/types/module"
	"github.com/mydexchain/chain-sdk/x/genutil"
	genutiltest "github.com/mydexchain/chain-sdk/x/genutil/client/testutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestAddGenesisAccountCmd(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	tests := []struct {
		name      string
		addr      string
		denom     string
		expectErr bool
	}{
		{
			name:      "invalid address",
			addr:      "",
			denom:     "1000atom",
			expectErr: true,
		},
		{
			name:      "valid address",
			addr:      addr1.String(),
			denom:     "1000atom",
			expectErr: false,
		},
		{
			name:      "multiple denoms",
			addr:      addr1.String(),
			denom:     "1000atom, 2000stake",
			expectErr: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			logger := log.NewNopLogger()
			cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
			require.NoError(t, err)

			appCodec, _ := simapp.MakeCodecs()
			err = genutiltest.ExecInitCmd(testMbm, home, appCodec)
			require.NoError(t, err)

			serverCtx := server.NewContext(viper.New(), cfg, logger)
			clientCtx := client.Context{}.WithJSONMarshaler(appCodec).WithHomeDir(home)

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

			cmd := simcmd.AddGenesisAccountCmd(home)
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
				fmt.Sprintf("--%s=home", flags.FlagHome)})

			if tc.expectErr {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}

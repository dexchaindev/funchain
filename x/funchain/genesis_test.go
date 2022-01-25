package funchain_test

import (
	"testing"

	keepertest "github.com/dexchaindev/funchain/testutil/keeper"
	"github.com/dexchaindev/funchain/testutil/nullify"
	"github.com/dexchaindev/funchain/x/funchain"
	"github.com/dexchaindev/funchain/x/funchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FunchainKeeper(t)
	funchain.InitGenesis(ctx, *k, genesisState)
	got := funchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package keeper_test

import (
	"testing"

	testkeeper "github.com/dexchaindev/funchain/testutil/keeper"
	"github.com/dexchaindev/funchain/x/funchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FunchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

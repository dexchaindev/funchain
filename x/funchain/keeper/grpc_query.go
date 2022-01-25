package keeper

import (
	"github.com/dexchaindev/funchain/x/funchain/types"
)

var _ types.QueryServer = Keeper{}

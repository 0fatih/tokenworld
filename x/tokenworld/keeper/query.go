package keeper

import (
	"tokenworld/x/tokenworld/types"
)

var _ types.QueryServer = Keeper{}

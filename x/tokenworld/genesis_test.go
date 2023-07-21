package tokenworld_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "tokenworld/testutil/keeper"
	"tokenworld/testutil/nullify"
	"tokenworld/x/tokenworld"
	"tokenworld/x/tokenworld/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenworldKeeper(t)
	tokenworld.InitGenesis(ctx, *k, genesisState)
	got := tokenworld.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

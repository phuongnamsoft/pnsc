package pns_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "pns/testutil/keeper"
	"pns/testutil/nullify"
	"pns/x/pns"
	"pns/x/pns/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PnsKeeper(t)
	pns.InitGenesis(ctx, *k, genesisState)
	got := pns.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

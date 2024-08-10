package cernwebfestpostquantumcryptography_test

import (
	"testing"

	keepertest "github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/testutil/keeper"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/testutil/nullify"
	cernwebfestpostquantumcryptography "github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/module"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CernwebfestpostquantumcryptographyKeeper(t)
	cernwebfestpostquantumcryptography.InitGenesis(ctx, k, genesisState)
	got := cernwebfestpostquantumcryptography.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

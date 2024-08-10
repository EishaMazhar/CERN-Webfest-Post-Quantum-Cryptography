package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/testutil/keeper"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CernwebfestpostquantumcryptographyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

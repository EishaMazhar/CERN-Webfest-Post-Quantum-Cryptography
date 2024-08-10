package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/testutil/keeper"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.CernwebfestpostquantumcryptographyKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

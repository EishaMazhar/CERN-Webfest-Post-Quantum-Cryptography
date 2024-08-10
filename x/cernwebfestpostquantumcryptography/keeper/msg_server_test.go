package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/testutil/keeper"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/keeper"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/x/cernwebfestpostquantumcryptography/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.CernwebfestpostquantumcryptographyKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

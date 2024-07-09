package mint_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tellor-io/layer/testutil/keeper"
	"github.com/tellor-io/layer/x/mint"
	"github.com/tellor-io/layer/x/mint/types"
)

func TestBeginBlocker(t *testing.T) {
	require := require.New(t)

	k, _, _, ctx := keeper.MintKeeper(t)
	err := mint.BeginBlocker(ctx, k)
	require.Error(err)

	// set minter
	minter := types.DefaultMinter()
	require.NoError(k.Minter.Set(ctx, minter))
	err = mint.BeginBlocker(ctx, k)
	require.Nil(err)

	// Initilize minter
	minter.Initialized = true
	require.NoError(k.Minter.Set(ctx, minter))
	err = mint.BeginBlocker(ctx, k)
	require.Nil(err)

	// advance time past 0
	ctx = ctx.WithBlockTime(time.Unix(1000, 0))
	err = mint.BeginBlocker(ctx, k)
	require.Nil(err)
}

func TestMintBlockProvision(t *testing.T) {
	require := require.New(t)

	k, _, bk, ctx := keeper.MintKeeper(t)
	minter := types.DefaultMinter()
	minter.Initialized = true

	// prev block time is 0
	require.NoError(k.Minter.Set(ctx, minter))
	err := mint.MintBlockProvision(ctx, k, time.Unix(1000, 0), minter)
	require.Nil(err)

	// prev block time 5 sec ago
	time5SecAgo := time.Now().Add(-5 * time.Second)
	minter.PreviousBlockTime = &time5SecAgo
	require.NoError(k.Minter.Set(ctx, minter))
	expectedAmt := math.NewInt(types.DailyMintRate * 5 * 1000 / types.MillisecondsInDay)
	bk.On("MintCoins", ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("loya", expectedAmt))).Return(nil)
	bk.On("InputOutputCoins", ctx, mock.Anything, mock.Anything).Return(nil)

	err = mint.MintBlockProvision(ctx, k, time.Now(), minter)
	require.Nil(err)
}

func TestSetPreviousBlockTime(t *testing.T) {
	require := require.New(t)

	k, _, _, ctx := keeper.MintKeeper(t)
	minter := types.DefaultMinter()
	minter.Initialized = true
	require.NoError(k.Minter.Set(ctx, minter))

	time1 := time.Unix(1000, 0)
	time2 := time.Unix(2000, 0)
	ctx = ctx.WithBlockTime(time1)
	err := mint.SetPreviousBlockTime(ctx, k, time2)
	require.Nil(err)

	minter, err = k.Minter.Get(ctx)
	require.Nil(err)
	require.Equal(minter.PreviousBlockTime.Unix(), time2.Unix())
}

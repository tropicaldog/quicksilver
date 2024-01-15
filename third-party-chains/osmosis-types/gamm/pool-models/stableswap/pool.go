package stableswap

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	sdkioerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/v7/third-party-chains/osmosis-types/gamm"
	"github.com/quicksilver-zone/quicksilver/v7/third-party-chains/osmosis-types/gamm/pool-models/internal/cfmm_common"
	"github.com/quicksilver-zone/quicksilver/v7/utils/addressutils"
)

var _ gamm.PoolI = &Pool{}

// NewStableswapPool returns a stableswap pool
// Invariants that are assumed to be satisfied and not checked:
// * len(initialLiquidity) = 2
// * FutureGovernor is valid
// * poolID doesn't already exist
func NewStableswapPool(poolId uint64, stableswapPoolParams PoolParams, initialLiquidity sdk.Coins, scalingFactors []uint64, futureGovernor string) (Pool, error) {
	if len(scalingFactors) == 0 {
		scalingFactors = []uint64{1, 1}
	} else if scalingFactors[0] == 0 || scalingFactors[1] == 0 {
		return Pool{}, gamm.ErrInvalidStableswapScalingFactors
	}

	pool := Pool{
		Address:            gamm.NewPoolAddress(poolId).String(),
		Id:                 poolId,
		PoolParams:         stableswapPoolParams,
		TotalShares:        sdk.NewCoin(gamm.GetPoolShareDenom(poolId), gamm.InitPoolSharesSupply),
		PoolLiquidity:      initialLiquidity,
		ScalingFactor:      scalingFactors,
		FuturePoolGovernor: futureGovernor,
	}

	return pool, nil
}

func (p Pool) GetAddress() sdk.AccAddress {
	addr, err := addressutils.AccAddressFromBech32(p.Address, "")
	if err != nil {
		panic(fmt.Sprintf("could not bech32 decode address of pool with id: %d", p.GetId()))
	}
	return addr
}

func (p Pool) String() string {
	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (p Pool) GetId() uint64 {
	return p.Id
}

func (p Pool) GetSwapFee(ctx sdk.Context) sdkmath.LegacyDec {
	return p.PoolParams.SwapFee
}

func (p Pool) GetExitFee(ctx sdk.Context) sdkmath.LegacyDec {
	return p.PoolParams.ExitFee
}

func (Pool) IsActive(ctx sdk.Context) bool {
	return true
}

// Returns the coins in the pool owned by all LP shareholders
func (p Pool) GetTotalPoolLiquidity(ctx sdk.Context) sdk.Coins {
	return p.PoolLiquidity
}

func (p Pool) GetTotalShares() sdkmath.Int {
	return p.TotalShares.Amount
}

func (p Pool) GetScalingFactors() []uint64 {
	return p.ScalingFactor
}

// CONTRACT: scaling factors follow the same index with pool liquidity denoms
func (p Pool) GetScalingFactorByLiquidityIndex(liquidityIndex int) uint64 {
	return p.ScalingFactor[liquidityIndex]
}

func (p Pool) NumAssets() int {
	return len(p.PoolLiquidity)
}

// getScaledPoolAmts returns scaled amount of pool liquidity based on each asset's precisions
func (p Pool) getScaledPoolAmts(denoms ...string) ([]sdkmath.LegacyDec, error) {
	result := make([]sdkmath.LegacyDec, len(denoms))
	poolLiquidity := p.PoolLiquidity
	liquidityIndexes := p.getLiquidityIndexMap()

	for i, denom := range denoms {
		liquidityIndex := liquidityIndexes[denom]

		amt := poolLiquidity.AmountOf(denom)
		if amt.IsZero() {
			return []sdkmath.LegacyDec{}, fmt.Errorf("denom %s does not exist in pool", denom)
		}
		scalingFactor := p.GetScalingFactorByLiquidityIndex(liquidityIndex)
		result[i] = sdkmath.LegacyNewDecFromInt(amt).QuoInt64Mut(int64(scalingFactor))
	}
	return result, nil
}

// getDescaledPoolAmts gets descaled amount of given denom and amount
func (p Pool) getDescaledPoolAmt(denom string, amount sdkmath.LegacyDec) sdkmath.LegacyDec {
	liquidityIndexes := p.getLiquidityIndexMap()
	liquidityIndex := liquidityIndexes[denom]

	scalingFactor := p.GetScalingFactorByLiquidityIndex(liquidityIndex)
	return amount.MulInt64(int64(scalingFactor))
}

// getLiquidityIndexMap creates a map of denoms to its index in pool liquidity
func (p Pool) getLiquidityIndexMap() map[string]int {
	poolLiquidity := p.PoolLiquidity
	liquidityIndexMap := make(map[string]int, poolLiquidity.Len())
	for i, coin := range poolLiquidity {
		liquidityIndexMap[coin.Denom] = i
	}
	return liquidityIndexMap
}

// updatePoolLiquidityForSwap updates the pool liquidity.
// It requires caller to validate that tokensIn and tokensOut only consist of
// denominations in the pool.
// The function sanity checks this, and panics if not the case.
func (p *Pool) updatePoolLiquidityForSwap(tokensIn sdk.Coins, tokensOut sdk.Coins) {
	numTokens := p.PoolLiquidity.Len()
	// update liquidity
	p.PoolLiquidity = p.PoolLiquidity.Add(tokensIn...).Sub(tokensOut...)
	// sanity check that no new denoms were added
	if len(p.PoolLiquidity) != numTokens {
		panic("updatePoolLiquidityForSwap changed number of tokens in pool")
	}
}

// updatePoolLiquidityForExit updates the pool liquidity after an exit.
// The function sanity checks that not all tokens of a given denom are removed,
// and panics if thats the case.
func (p *Pool) updatePoolLiquidityForExit(tokensOut sdk.Coins) {
	p.updatePoolLiquidityForSwap(sdk.Coins{}, tokensOut)
}

func (p *Pool) updatePoolForJoin(tokensIn sdk.Coins, newShares sdkmath.Int) {
	numTokens := p.NumAssets()
	p.PoolLiquidity = p.PoolLiquidity.Add(tokensIn...)
	if len(p.PoolLiquidity) != numTokens {
		panic(fmt.Sprintf("updatePoolForJoin changed number of tokens in pool from %d to %d", numTokens, len(p.PoolLiquidity)))
	}
	p.TotalShares.Amount = p.TotalShares.Amount.Add(newShares)
}

// TODO: These should all get moved to amm.go
func (p Pool) CalcOutAmtGivenIn(ctx sdk.Context, tokenIn sdk.Coins, tokenOutDenom string, swapFee sdkmath.LegacyDec) (tokenOut sdk.Coin, err error) {
	if tokenIn.Len() != 1 {
		return sdk.Coin{}, errors.New("stableswap CalcOutAmtGivenIn: tokenIn is of wrong length")
	}
	outAmtDec, err := p.calcOutAmtGivenIn(tokenIn[0], tokenOutDenom, swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}

	// we ignore the decimal component, as token out amount must round down
	tokenOutAmt := outAmtDec.TruncateInt()
	if !tokenOutAmt.IsPositive() {
		return sdk.Coin{}, sdkioerrors.Wrapf(gamm.ErrInvalidMathApprox, "token amount must be positive")
	}
	return sdk.NewCoin(tokenOutDenom, tokenOutAmt), nil
}

func (p *Pool) SwapOutAmtGivenIn(ctx sdk.Context, tokenIn sdk.Coins, tokenOutDenom string, swapFee sdkmath.LegacyDec) (tokenOut sdk.Coin, err error) {
	tokenOut, err = p.CalcOutAmtGivenIn(ctx, tokenIn, tokenOutDenom, swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}

	p.updatePoolLiquidityForSwap(tokenIn, sdk.NewCoins(tokenOut))

	return tokenOut, nil
}

func (p Pool) CalcInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coins, tokenInDenom string, swapFee sdkmath.LegacyDec) (tokenIn sdk.Coin, err error) {
	if tokenOut.Len() != 1 {
		return sdk.Coin{}, errors.New("stableswap CalcInAmtGivenOut: tokenOut is of wrong length")
	}
	// TODO: Refactor this later to handle scaling factors
	amt, err := p.calcInAmtGivenOut(tokenOut[0], tokenInDenom, swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}

	// We round up tokenInAmt, as this is whats charged for the swap, for the precise amount out.
	// Otherwise, the pool would under-charge by this rounding error.
	tokenInAmt := amt.Ceil().TruncateInt()

	if !tokenInAmt.IsPositive() {
		return sdk.Coin{}, sdkioerrors.Wrapf(gamm.ErrInvalidMathApprox, "token amount must be positive")
	}
	return sdk.NewCoin(tokenInDenom, tokenInAmt), nil
}

func (p *Pool) SwapInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coins, tokenInDenom string, swapFee sdkmath.LegacyDec) (tokenIn sdk.Coin, err error) {
	tokenIn, err = p.CalcInAmtGivenOut(ctx, tokenOut, tokenInDenom, swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}

	p.updatePoolLiquidityForSwap(sdk.NewCoins(tokenIn), tokenOut)

	return tokenIn, nil
}

func (p Pool) SpotPrice(ctx sdk.Context, baseAssetDenom string, quoteAssetDenom string) (sdkmath.LegacyDec, error) {
	reserves, err := p.getScaledPoolAmts(baseAssetDenom, quoteAssetDenom)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	scaledSpotPrice := spotPrice(reserves[0], reserves[1])
	spotPrice := p.getDescaledPoolAmt(baseAssetDenom, scaledSpotPrice)

	return spotPrice, nil
}

func (p Pool) Copy() Pool {
	p2 := p
	p2.PoolLiquidity = sdk.NewCoins(p.PoolLiquidity...)
	return p2
}

func (p *Pool) CalcJoinPoolShares(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdkmath.LegacyDec) (numShares sdkmath.Int, newLiquidity sdk.Coins, err error) {
	pCopy := p.Copy()
	return pCopy.joinPoolSharesInternal(ctx, tokensIn, swapFee)
}

// TODO: implement this
func (*Pool) CalcJoinPoolNoSwapShares(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdkmath.LegacyDec) (numShares sdkmath.Int, newLiquidity sdk.Coins, err error) {
	return sdkmath.ZeroInt(), nil, err
}

func (p *Pool) JoinPool(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdkmath.LegacyDec) (numShares sdkmath.Int, err error) {
	numShares, _, err = p.joinPoolSharesInternal(ctx, tokensIn, swapFee)
	return numShares, err
}

// TODO: implement this
func (*Pool) JoinPoolNoSwap(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdkmath.LegacyDec) (numShares sdkmath.Int, err error) {
	return sdkmath.ZeroInt(), err
}

func (p *Pool) ExitPool(ctx sdk.Context, exitingShares sdkmath.Int, exitFee sdkmath.LegacyDec) (exitingCoins sdk.Coins, err error) {
	exitingCoins, err = p.CalcExitPoolCoinsFromShares(ctx, exitingShares, exitFee)
	if err != nil {
		return sdk.Coins{}, err
	}

	p.TotalShares.Amount = p.TotalShares.Amount.Sub(exitingShares)
	p.updatePoolLiquidityForExit(exitingCoins)

	return exitingCoins, nil
}

func (p Pool) CalcExitPoolCoinsFromShares(ctx sdk.Context, exitingShares sdkmath.Int, exitFee sdkmath.LegacyDec) (exitingCoins sdk.Coins, err error) {
	return cfmm_common.CalcExitPool(ctx, &p, exitingShares, exitFee)
}

// no-op for stableswap
func (*Pool) PokePool(blockTime time.Time) {}

// SetStableSwapScalingFactors sets scaling factors for pool to the given amount
// It should only be able to be successfully called by the pool's ScalingFactorGovernor
// TODO: move commented test for this function from x/gamm/keeper/pool_service_test.go once a pool_test.go file has been created for stableswap
func (p *Pool) SetStableSwapScalingFactors(ctx sdk.Context, scalingFactors []uint64, scalingFactorGovernor string) error {
	if scalingFactorGovernor != p.ScalingFactorGovernor {
		return gamm.ErrNotScalingFactorGovernor
	}

	if len(scalingFactors) != p.PoolLiquidity.Len() {
		return gamm.ErrInvalidStableswapScalingFactors
	}

	p.ScalingFactor = scalingFactors

	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)

	HasSupply(ctx sdk.Context, denom string) bool

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error

	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin

	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type AccountKeeper interface {
	SetModuleAccount(ctx sdk.Context, macc sdk.ModuleAccountI)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI
}

// CommunityPoolKeeper defines the contract needed to be fulfilled for community pool interactions.
type CommunityPoolKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

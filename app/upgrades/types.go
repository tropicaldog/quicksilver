package upgrades

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/quicksilver-zone/quicksilver/app/keepers"
)

// upgrade name consts: vMMmmppUpgradeName (M=Major, m=minor, p=patch).
const (
	ProductionChainID      = "quicksilver-2"
	RhyeChainID            = "rhye-2"
	DevnetChainID          = "magic-2"
	TestChainID            = "testchain1"
	OsmosisTestnetChainID  = "osmo-test-5"
	JunoTestnetChainID     = "uni-6"
	StargazeTestnetChainID = "elgafar-1"
	SommelierChainID       = "sommelier-3"

	// testnet upgrades
	V010402rc1UpgradeName    = "v1.4.2-rc1"
	V010402rc2UpgradeName    = "v1.4.2-rc2"
	V010402rc3UpgradeName    = "v1.4.2-rc3"
	V010402rc4UpgradeName    = "v1.4.2-rc4"
	V010402rc5UpgradeName    = "v1.4.2-rc5"
	V010402rc6UpgradeName    = "v1.4.2-rc6"
	V010402rc7UpgradeName    = "v1.4.2-rc7"
	V010403rc0UpgradeName    = "v1.4.3-rc0"
	V010404beta0UpgradeName  = "v1.4.4-beta.0"
	V010404beta1UpgradeName  = "v1.4.4-beta.1"
	V010404beta5UpgradeName  = "v1.4.4-beta.5"
	V010404beta7UpgradeName  = "v1.4.4-beta.7"
	V010404rc0UpgradeName    = "v1.4.4-rc.0"
	V010404beta8UpgradeName  = "v1.4.4-beta.8"
	V010404rc1UpgradeName    = "v1.4.4-rc.1"
	V010404beta9UpgradeName  = "v1.4.4-beta.9"
	V010404beta10UpgradeName = "v1.4.4-beta.10"
	V010404rc2UpgradeName    = "v1.4.4-rc.2"
	V010404rc3UpgradeName    = "v1.4.4-rc.3"
	V010404rc4UpgradeName    = "v1.4.4-rc.4"
	V010405rc0UpgradeName    = "v1.4.5-rc0"
	V010405rc2UpgradeName    = "v1.4.5-rc2"
	V010405rc3UpgradeName    = "v1.4.5-rc3"
	V010405rc4UpgradeName    = "v1.4.5-rc4"

	// mainnet upgrades
	V010405UpgradeName = "v1.4.5"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(*module.Manager, module.Configurator, *keepers.AppKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades storetypes.StoreUpgrades
}

// nolint:all //function useful for writing network specific upgrade handlers
func isTest(ctx sdk.Context) bool {
	return ctx.ChainID() == TestChainID
}

// nolint:all //function useful for writing network specific upgrade handlers
func isDevnet(ctx sdk.Context) bool {
	return ctx.ChainID() == DevnetChainID
}

// nolint:all //function useful for writing network specific upgrade handlers
func isTestnet(ctx sdk.Context) bool {
	return ctx.ChainID() == RhyeChainID
}

// nolint:all //function useful for writing network specific upgrade handlers
func isMainnet(ctx sdk.Context) bool {
	return ctx.ChainID() == ProductionChainID
}

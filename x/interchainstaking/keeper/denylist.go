package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/utils"
	"github.com/quicksilver-zone/quicksilver/x/interchainstaking/types"
)

// unbondigng records are keyed by chainId, validator and epoch, as they must be unique with regard to this triple.
func GetDenyListKey(chainID string, address sdk.ValAddress) []byte {
	return append(types.KeyPrefixDenyList, append([]byte(chainID), address.Bytes()...)...)
}

// HasDenyListRecord returns whether a deny list record for the given chainID and val address exists in the KV store.
func (k Keeper) HasDenyListRecord(ctx sdk.Context, chainID string, valAddr sdk.ValAddress) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := store.Get(GetDenyListKey(chainID, valAddr))
	if bz == nil {
		return false
	}

	return true
}

// SetDenyListRecord stores the deny list flag.
func (k Keeper) SetDenyListRecord(ctx sdk.Context, chainID string, valAddr sdk.ValAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := []byte{0x01}
	store.Set(GetDenyListKey(chainID, valAddr), bz)
}

// DeleteDenyListRecord deletes deny list record
func (k Keeper) DeleteDenyListRecord(ctx sdk.Context, chainID string, valAddr sdk.ValAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	store.Delete(GetDenyListKey(chainID, valAddr))
}

// IteratePrefixedRedelegationRecords iterate through all records with given prefix
func (k Keeper) IteratePrefixedDenyListRecords(ctx sdk.Context, prefixBytes []byte, fn func(index int64, key []byte) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRedelegationRecord)

	iterator := sdk.KVStorePrefixIterator(store, prefixBytes)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {

		stop := fn(i, iterator.Key())

		if stop {
			break
		}
		i++
	}
}

// IterateDenyListRecords iterate through all records
func (k Keeper) IterateDenyListRecords(ctx sdk.Context, fn func(index int64, key []byte) (stop bool)) {
	k.IteratePrefixedDenyListRecords(ctx, nil, fn)
}

// IterateDenyListRecordsForChain iterate through all records for a given chainID
func (k Keeper) IterateDenyListRecordsForChain(ctx sdk.Context, chainID string, fn func(index int64, key []byte) (stop bool)) {
	k.IteratePrefixedDenyListRecords(ctx, []byte(chainID), fn)
}

// Helper functions

// Convenience function to check deny status by valoper string.
func (k Keeper) IsDenyListed(ctx sdk.Context, chainID string, valoper string) bool {
	valAddr, err := utils.ValAddressFromBech32(valoper, "")
	if err != nil {
		return false
	}
	return k.HasDenyListRecord(ctx, chainID, valAddr)
}

func (k Keeper) AddDenyListEntry(ctx sdk.Context, chainID string, valoper string) error {
	valAddr, err := utils.ValAddressFromBech32(valoper, "")
	if err != nil {
		return err
	}
	k.SetDenyListRecord(ctx, chainID, valAddr)
	return nil
}

func (k Keeper) RemoveDenyListEntry(ctx sdk.Context, chainID string, valoper string) error {
	valAddr, err := utils.ValAddressFromBech32(valoper, "")
	if err != nil {
		return err
	}
	k.DeleteDenyListRecord(ctx, chainID, valAddr)
	return nil
}

package v043

import (
	sdk "github.com/puneetsingh166/tm-load-test/types"
	v043distribution "github.com/puneetsingh166/tm-load-test/x/distribution/migrations/v043"
	v040slashing "github.com/puneetsingh166/tm-load-test/x/slashing/migrations/v040"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v040slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v040slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v040slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}

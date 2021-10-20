package v044

import (
	"github.com/puneetsingh166/tm-load-test/codec"
	"github.com/puneetsingh166/tm-load-test/store/prefix"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/types/address"
	v043 "github.com/puneetsingh166/tm-load-test/x/bank/migrations/v043"
	"github.com/puneetsingh166/tm-load-test/x/bank/types"
)

// MigrateStore performs in-place store migrations from v0.43 to v0.44. The
// migration includes:
//
// - Migrate coin storage to save only amount.
// - Add an additional reverse index from denomination to address.
// - Remove duplicate denom from denom metadata store key.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	err := addDenomReverseIndex(store, cdc)
	if err != nil {
		return err
	}

	return migrateDenomMetadata(store)
}

func addDenomReverseIndex(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldBalancesStore := prefix.NewStore(store, v043.BalancesPrefix)

	oldBalancesIter := oldBalancesStore.Iterator(nil, nil)
	defer oldBalancesIter.Close()

	denomPrefixStores := make(map[string]prefix.Store) // memoize prefix stores

	for ; oldBalancesIter.Valid(); oldBalancesIter.Next() {
		var balance sdk.Coin
		if err := cdc.Unmarshal(oldBalancesIter.Value(), &balance); err != nil {
			return err
		}

		addr, err := v043.AddressFromBalancesStore(oldBalancesIter.Key())
		if err != nil {
			return err
		}

		var coin sdk.DecCoin
		if err := cdc.Unmarshal(oldBalancesIter.Value(), &coin); err != nil {
			return err
		}

		bz, err := coin.Amount.Marshal()
		if err != nil {
			return err
		}

		newStore := prefix.NewStore(store, types.CreateAccountBalancesPrefix(addr))
		newStore.Set([]byte(coin.Denom), bz)

		denomPrefixStore, ok := denomPrefixStores[balance.Denom]
		if !ok {
			denomPrefixStore = prefix.NewStore(store, CreateAddressDenomPrefix(balance.Denom))
			denomPrefixStores[balance.Denom] = denomPrefixStore
		}

		// Store a reverse index from denomination to account address with a
		// sentinel value.
		denomPrefixStore.Set(address.MustLengthPrefix(addr), []byte{0})
	}

	return nil
}

func migrateDenomMetadata(store sdk.KVStore) error {
	oldDenomMetaDataStore := prefix.NewStore(store, v043.DenomMetadataPrefix)

	oldDenomMetaDataIter := oldDenomMetaDataStore.Iterator(nil, nil)
	defer oldDenomMetaDataIter.Close()

	for ; oldDenomMetaDataIter.Valid(); oldDenomMetaDataIter.Next() {
		oldKey := oldDenomMetaDataIter.Key()
		// old key: prefix_bytes | denom_bytes | denom_bytes
		newKey := append(types.DenomMetadataPrefix, oldKey[:len(oldKey)/2+1]...)

		store.Set(newKey, oldDenomMetaDataIter.Value())
		oldDenomMetaDataStore.Delete(oldDenomMetaDataIter.Key())
	}

	return nil
}

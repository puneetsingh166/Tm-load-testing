package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/puneetsingh166/tm-load-test/store/cache"
	"github.com/puneetsingh166/tm-load-test/store/rootmulti"
	"github.com/puneetsingh166/tm-load-test/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}

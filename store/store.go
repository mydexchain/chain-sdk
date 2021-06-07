package store

import (
	dbm "github.com/mydexchain/tm-db"

	"github.com/mydexchain/chain-sdk/store/cache"
	"github.com/mydexchain/chain-sdk/store/rootmulti"
	"github.com/mydexchain/chain-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}

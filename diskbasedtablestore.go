package markov

import (
	"github.com/mandykoh/keva"
)

type DiskBasedTableStore struct {
	KevaStore *keva.Store
}

func (ts *DiskBasedTableStore) Close() error {
	return ts.KevaStore.Close()
}

func (ts *DiskBasedTableStore) Get(key SequenceKey, dest *Table) error {
	return ts.KevaStore.Get(string(key), dest)
}

func (ts *DiskBasedTableStore) Put(key SequenceKey, table *Table) error {
	return ts.KevaStore.Put(string(key), table)
}

func NewDiskBasedTableStore(rootPath string) *DiskBasedTableStore {
	kevaStore := keva.NewStore(rootPath)

	return &DiskBasedTableStore{
		KevaStore: kevaStore,
	}
}

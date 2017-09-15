package markov

import (
	"github.com/mandykoh/keva"
)

type KevaTableStore struct {
	KevaStore *keva.Store
}

func (ts *KevaTableStore) Close() error {
	return ts.KevaStore.Close()
}

func (ts *KevaTableStore) Get(key SequenceKey, dest *Table) error {
	err := ts.KevaStore.Get(string(key), dest)
	if err == keva.ErrValueNotFound {
		return ErrTableNotFound
	}
	return err
}

func (ts *KevaTableStore) Put(key SequenceKey, table *Table) error {
	return ts.KevaStore.Put(string(key), table)
}

func NewKevaTableStore(rootPath string) (*KevaTableStore, error) {
	kevaStore, err := keva.NewStore(rootPath)
	if err != nil {
		return nil, err
	}

	return &KevaTableStore{
		KevaStore: kevaStore,
	}, nil
}

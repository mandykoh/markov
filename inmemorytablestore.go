package markov

type InMemoryTableStore struct {
	TablesBySequence map[SequenceKey]*Table
}

func (ts *InMemoryTableStore) Get(key SequenceKey, dest *Table) error {
	t, ok := ts.TablesBySequence[key]
	if !ok {
		return ErrTableNotFound
	}

	*dest = *t

	return nil
}

func (ts *InMemoryTableStore) Put(key SequenceKey, table *Table) error {
	ts.TablesBySequence[key] = table
	return nil
}

func NewInMemoryTableStore() *InMemoryTableStore {
	return &InMemoryTableStore{
		TablesBySequence: make(map[SequenceKey]*Table),
	}
}

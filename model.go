package markov

type Model struct {
	Tables TableStore
}

func (m *Model) Add(seq Sequence, nextSymbol string) error {
	seqKey := seq.Key()

	var t Table
	err := m.Tables.Get(seqKey, &t)

	if err == ErrTableNotFound {
		t = EmptyTable()
	} else if err != nil {
		return err
	}

	t.Add(nextSymbol)
	return m.Tables.Put(seqKey, &t)
}

func (m Model) Sample(seq Sequence, sampleValue float64) (nextSymbol string, err error) {
	seqKey := seq.Key()

	var t Table
	err = m.Tables.Get(seqKey, &t)
	if err == ErrTableNotFound {
		nextSymbol = ""
		err = nil
		return

	} else if err != nil {
		return
	}

	symbolIndex := uint64(sampleValue * float64(t.TotalSymbols))
	nextSymbol = t.Sample(symbolIndex)
	return
}

func NewModel(store TableStore) *Model {
	return &Model{
		Tables: store,
	}
}

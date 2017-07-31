package markov

type Table struct {
	Entries        []TableEntry
	EntryIndices   map[SymbolKey]int
	TotalFrequency uint64
}

func (t *Table) Add(s Symbol) {
	key := s.Key()

	index, exists := t.EntryIndices[key]
	if !exists {
		index = len(t.Entries)
		t.EntryIndices[key] = index

		entry := TableEntry{
			Frequency: 0,
			Symbol:    s,
		}
		t.Entries = append(t.Entries, entry)
	}

	t.Entries[index].Frequency++
	t.TotalFrequency++
}

func EmptyTable() Table {
	return Table{
		EntryIndices: make(map[SymbolKey]int),
	}
}

package markov

type Table struct {
	Entries      []TableEntry
	EntryIndices map[SymbolKey]int
}

func (t *Table) Add(s Symbol) {
	key := s.Key()

	index, exists := t.EntryIndices[key]
	if !exists {
		t.EntryIndices[key] = len(t.Entries)

		entry := TableEntry{
			Frequency: 1,
			Symbol:    s,
		}
		t.Entries = append(t.Entries, entry)
		return
	}

	t.Entries[index].Frequency++
}

func NewTable() *Table {
	return &Table{
		EntryIndices: make(map[SymbolKey]int),
	}
}

type TableEntry struct {
	Frequency uint64
	Symbol    Symbol
}

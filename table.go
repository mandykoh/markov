package markov

type Table struct {
	TotalSymbols uint64
	Entries      []TableEntry
	EntryIndices map[SymbolKey]int
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
	t.TotalSymbols++

	t.sortEntry(index)
}

func (t Table) Sample(symbolIndex uint64) Symbol {
	remaining := symbolIndex

	for index := 0; index < len(t.Entries); index++ {
		entry := &t.Entries[index]

		if remaining < entry.Frequency {
			return entry.Symbol
		}
		remaining -= entry.Frequency
	}

	return nil
}

func (t *Table) sortEntry(index int) {
	j := index

	for i := index - 1; i >= 0 && t.Entries[j].Frequency > t.Entries[i].Frequency; i-- {
		tmp := t.Entries[i]
		t.Entries[i] = t.Entries[j]
		t.Entries[j] = tmp

		j = i
	}

	for i := j; i <= index; i++ {
		t.EntryIndices[t.Entries[i].Symbol.Key()] = i
	}
}

func EmptyTable() Table {
	return Table{
		EntryIndices: make(map[SymbolKey]int),
	}
}

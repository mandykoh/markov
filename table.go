package markov

type Table struct {
	TotalSymbols uint64
	Entries      []TableEntry
	EntryIndices map[string]int
}

func (t *Table) Add(symbol string) {
	index, exists := t.EntryIndices[symbol]
	if !exists {
		index = len(t.Entries)
		t.EntryIndices[symbol] = index

		entry := TableEntry{
			Frequency: 0,
			Symbol:    symbol,
		}
		t.Entries = append(t.Entries, entry)
	}

	t.Entries[index].Frequency++
	t.TotalSymbols++

	t.sortEntry(index)
}

func (t Table) Sample(symbolIndex uint64) (symbol string) {
	remaining := symbolIndex

	for index := 0; index < len(t.Entries); index++ {
		entry := &t.Entries[index]

		if remaining < entry.Frequency {
			return entry.Symbol
		}
		remaining -= entry.Frequency
	}

	return ""
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
		t.EntryIndices[t.Entries[i].Symbol] = i
	}
}

func EmptyTable() Table {
	return Table{
		EntryIndices: make(map[string]int),
	}
}

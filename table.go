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

	t.sortEntry(index)
}

func (t *Table) sortEntry(index int) {
	j := index

	for i := index - 1; i >= 0 && t.Entries[j].Frequency > t.Entries[i].Frequency; i-- {
		t.EntryIndices[t.Entries[i].Symbol.Key()] = j
		t.EntryIndices[t.Entries[j].Symbol.Key()] = i

		tmp := t.Entries[i]
		t.Entries[i] = t.Entries[j]
		t.Entries[j] = tmp

		j = i
	}
}

func EmptyTable() Table {
	return Table{
		EntryIndices: make(map[SymbolKey]int),
	}
}

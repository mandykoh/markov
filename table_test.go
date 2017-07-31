package markov

import (
	"testing"
)

func TestTable(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		t.Run("creates an entry for new symbols", func(t *testing.T) {
			table := EmptyTable()
			table.Add(StringSymbol("a"))

			indexA := table.EntryIndices[SymbolKey("a")]

			if actual := len(table.Entries); actual != 1 {
				t.Fatalf("Expected 1 entry but got %d", actual)
			}
			if expected, actual := uint64(1), table.Entries[indexA].Frequency; expected != actual {
				t.Fatalf("Expected frequency of %d but got %d", expected, actual)
			}
			if expected, actual := uint64(1), table.TotalFrequency; expected != actual {
				t.Fatalf("Expected total frequency of %d but got %d", expected, actual)
			}

			table.Add(StringSymbol("b"))
			indexB := table.EntryIndices[SymbolKey("b")]

			if actual := len(table.Entries); actual != 2 {
				t.Fatalf("Expected 2 entries but got %d", actual)
			}
			if expected, actual := uint64(1), table.Entries[indexB].Frequency; expected != actual {
				t.Fatalf("Expected frequency of %d but got %d", expected, actual)
			}
			if expected, actual := uint64(2), table.TotalFrequency; expected != actual {
				t.Fatalf("Expected total frequency of %d but got %d", expected, actual)
			}
		})

		t.Run("increments frequency for existing symbols", func(t *testing.T) {
			table := EmptyTable()
			table.Add(StringSymbol("a"))
			table.Add(StringSymbol("a"))

			if actual := len(table.Entries); actual != 1 {
				t.Fatalf("Expected 1 entries but got %d", actual)
			}
			if expected, actual := uint64(2), table.Entries[0].Frequency; expected != actual {
				t.Fatalf("Expected frequency of %d but got %d", expected, actual)
			}
			if expected, actual := uint64(2), table.TotalFrequency; expected != actual {
				t.Fatalf("Expected total frequency of %d but got %d", expected, actual)
			}
		})

		t.Run("keeps entries sorted from most to least frequent", func(t *testing.T) {
			table := EmptyTable()
			table.Add(StringSymbol("a"))
			table.Add(StringSymbol("b"))
			table.Add(StringSymbol("c"))
			table.Add(StringSymbol("b"))
			table.Add(StringSymbol("d"))
			table.Add(StringSymbol("d"))
			table.Add(StringSymbol("d"))

			scenarios := []struct {
				ExpectedSymbol    Symbol
				ExpectedFrequency uint64
			}{
				{StringSymbol("d"), uint64(3)},
				{StringSymbol("b"), uint64(2)},
				{StringSymbol("a"), uint64(1)},
				{StringSymbol("c"), uint64(1)},
			}

			if expected, actual := len(scenarios), len(table.Entries); expected != actual {
				t.Fatalf("Expected %d entries but got %d", expected, actual)
			}

			for i, scenario := range scenarios {
				if actual := table.Entries[i].Symbol; scenario.ExpectedSymbol != actual {
					t.Errorf("Expected symbol '%v' in position %d but got '%v'", scenario.ExpectedSymbol, i, actual)
				}
				if actual := table.Entries[i].Frequency; scenario.ExpectedFrequency != actual {
					t.Errorf("Expected frequency of %d in position %d but got %d", scenario.ExpectedFrequency, i, actual)
				}
			}
		})
	})
}

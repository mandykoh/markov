package markov

import "testing"

func TestTable(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		t.Run("creates an entry for new symbols", func(t *testing.T) {
			table := EmptyTable()
			table.Add(StringSymbol("a"))

			if actual := len(table.Entries); actual != 1 {
				t.Fatalf("Expected 1 entry but got %d", actual)
			}
			if expected, actual := uint64(1), table.Entries[0].Frequency; expected != actual {
				t.Fatalf("Expected frequency of %d but got %d", expected, actual)
			}

			table.Add(StringSymbol("b"))

			if actual := len(table.Entries); actual != 2 {
				t.Fatalf("Expected 2 entries but got %d", actual)
			}
			if expected, actual := uint64(1), table.Entries[1].Frequency; expected != actual {
				t.Fatalf("Expected frequency of %d but got %d", expected, actual)
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
		})
	})
}

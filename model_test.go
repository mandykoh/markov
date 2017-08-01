package markov

import "testing"

func TestModel(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		t.Run("creates a mapping for a new sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)
			s1 := EmptySequence(2).WithNext("a")
			s2 := EmptySequence(2).WithNext("b")

			m.Add(s1, "b")
			m.Add(s2, "b")

			if expected, actual := 2, len(ts.TablesBySequence); expected != actual {
				t.Fatalf("Expected %d tables but found %d", expected, actual)
			}
			if expected, actual := uint64(1), ts.TablesBySequence[s1.Key()].TotalSymbols; expected != actual {
				t.Fatalf("Expected one symbol for sequence '%s' but found %d", s1.Key(), actual)
			}
			if expected, actual := uint64(1), ts.TablesBySequence[s2.Key()].TotalSymbols; expected != actual {
				t.Fatalf("Expected one symbol for sequence '%s' but found %d", s2.Key(), actual)
			}
		})

		t.Run("adds to a mapping for an existing sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)
			s := EmptySequence(2).WithNext("a")

			m.Add(s, "b")
			m.Add(s, "c")

			if expected, actual := 1, len(ts.TablesBySequence); expected != actual {
				t.Fatalf("Expected one table but found %d", actual)
			}
			if expected, actual := uint64(2), ts.TablesBySequence[s.Key()].TotalSymbols; expected != actual {
				t.Fatalf("Expected %d symbols but found %d", expected, actual)
			}
		})
	})

	t.Run("Sample()", func(t *testing.T) {

		t.Run("samples the table mapped to the sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)
			s := EmptySequence(2).WithNext("a")

			m.Add(s, "b")

			nextSymbol, err := m.Sample(s, 0.0)

			if err != nil {
				t.Fatalf("Error sampling from model: %v", err)
			}
			if expected, actual := "b", nextSymbol; expected != actual {
				t.Errorf("Expected symbol '%v' but got '%v'", expected, actual)
			}
		})

		t.Run("returns an error if no table is mapped to the sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)
			s := EmptySequence(2)

			_, err := m.Sample(s, 0)

			if err != nil {
				if err != ErrTableNotFound {
					t.Fatalf("Expected ErrTableNotFound but got %v", err)
				}
			}
		})
	})
}

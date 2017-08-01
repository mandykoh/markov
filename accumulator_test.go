package markov

import "testing"

func TestAccumulator(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		t.Run("adds a symbol to a table by rotating the sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)
			acc := NewAccumulator(m, 2)

			err := acc.Add("a")
			if err != nil {
				t.Fatalf("Error accumulating: %v", err)
			}

			if expected, actual := 1, len(ts.TablesBySequence); expected != actual {
				t.Errorf("Expected one table but found %d", actual)
			}
			if expected, actual := SequenceKey("a"), acc.CurrentSequence.Key(); expected != actual {
				t.Errorf("Expected sequence to be '%s' but was '%s'", expected, actual)
			}

			err = acc.Add("b")
			if err != nil {
				t.Fatalf("Error accumulating: %v", err)
			}

			if expected, actual := 2, len(ts.TablesBySequence); expected != actual {
				t.Errorf("Expected %d tables but found %d", expected, actual)
			}
			if expected, actual := SequenceKey("a|b"), acc.CurrentSequence.Key(); expected != actual {
				t.Errorf("Expected sequence to be '%s' but was '%s'", expected, actual)
			}

			err = acc.Add("c")
			if err != nil {
				t.Fatalf("Error accumulating: %v", err)
			}

			if expected, actual := 3, len(ts.TablesBySequence); expected != actual {
				t.Errorf("Expected %d tables but found %d", expected, actual)
			}
			if expected, actual := SequenceKey("b|c"), acc.CurrentSequence.Key(); expected != actual {
				t.Errorf("Expected sequence to be '%s' but was '%s'", expected, actual)
			}
		})
	})
}

package markov

import "testing"

func TestSequence(t *testing.T) {

	expectSequenceMatches := func(t *testing.T, s Sequence, symbols ...string) {
		if sequenceOrder, symbolCount := s.Order(), uint(len(symbols)); sequenceOrder != symbolCount {
			t.Fatalf("Expected sequence of order %d to match %d symbols", sequenceOrder, symbolCount)
		}

		for i := 0; i < len(s.Symbols); i++ {
			if symbols[i] == "" {
				if s.Symbols[i] != "" {
					t.Errorf("Expected symbol %d to be empty but was '%v'", i, s.Symbols[i])
				}
			} else if symbols[i] != s.Symbols[i] {
				t.Errorf("Expected symbol %d to be '%v' but was '%v'", i, symbols[i], s.Symbols[i])
			}
		}
	}

	t.Run("EmptySequence()", func(t *testing.T) {

		t.Run("creates Sequence of the specified order", func(t *testing.T) {
			s := EmptySequence(3)

			if expected, actual := 3, len(s.Symbols); actual != expected {
				t.Fatalf("Expected %d symbols but found %d", expected, actual)
			}

			for i := 0; i < len(s.Symbols); i++ {
				if s.Symbols[i] != "" {
					t.Errorf("Expected symbol %d in sequence to be empty but got '%v'", i, s.Symbols[i])
				}
			}
		})
	})

	t.Run("Key()", func(t *testing.T) {

		t.Run("returns a sequence key derived from the symbols", func(t *testing.T) {
			s := SequenceWith("a", "b", "c")

			if expected, actual := SequenceKey("a|b|c"), s.Key(); expected != actual {
				t.Errorf("Expected sequence key '%s' but got '%s'", expected, actual)
			}
		})

		t.Run("correctly returns keys that have leading empty symbols", func(t *testing.T) {
			s := SequenceWith("", "a", "b")

			if expected, actual := SequenceKey("a|b"), s.Key(); expected != actual {
				t.Errorf("Expected sequence key '%s' but got '%s'", expected, actual)
			}
		})

		t.Run("correctly treats non-leading empty symbols as significant", func(t *testing.T) {
			s := SequenceWith("a", "", "b")

			if expected, actual := SequenceKey("a||b"), s.Key(); expected != actual {
				t.Errorf("Expected sequence key '%s' but got '%s'", expected, actual)
			}
		})

		t.Run("correctly escapes the separator", func(t *testing.T) {
			s := SequenceWith(`a|b\`, "c")

			if expected, actual := SequenceKey(`a\|b\\|c`), s.Key(); expected != actual {
				t.Errorf("Expected sequence key '%s' but got '%s'", expected, actual)
			}
		})
	})

	t.Run("WithNext()", func(t *testing.T) {

		t.Run("returns a sequence using the specified next symbol", func(t *testing.T) {
			s := EmptySequence(3)

			s = s.WithNext("a")
			expectSequenceMatches(t, s, "", "", "a")

			s = s.WithNext("b")
			expectSequenceMatches(t, s, "", "a", "b")

			s = s.WithNext("c")
			expectSequenceMatches(t, s, "a", "b", "c")

			s = s.WithNext("d")
			expectSequenceMatches(t, s, "b", "c", "d")
		})
	})
}

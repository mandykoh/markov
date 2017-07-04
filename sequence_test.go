package markov

import "testing"

func TestSequence(t *testing.T) {

	expectSequenceMatches := func(t *testing.T, s Sequence, symbols ...Symbol) {
		if sequenceOrder, symbolCount := s.Order(), uint(len(symbols)); sequenceOrder != symbolCount {
			t.Fatalf("Expected sequence of order %d to match %d symbols", sequenceOrder, symbolCount)
		}

		for i := 0; i < len(s.Symbols); i++ {
			if symbols[i] == nil {
				if s.Symbols[i] != nil {
					t.Errorf("Expected symbol %d to be nil but was '%v'", i, s.Symbols[i])
				}
			} else if !symbols[i].Equals(s.Symbols[i]) {
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
				if s.Symbols[i] != nil {
					t.Errorf("Expected symbol %d in sequence to be nil but got '%v'", i, s.Symbols[i])
				}
			}
		})
	})

	t.Run("Next()", func(t *testing.T) {

		t.Run("returns next sequence using the specified next symbol", func(t *testing.T) {
			s := EmptySequence(3)

			s = s.Next(StringSymbol("a"))
			expectSequenceMatches(t, s, nil, nil, StringSymbol("a"))

			s = s.Next(StringSymbol("b"))
			expectSequenceMatches(t, s, nil, StringSymbol("a"), StringSymbol("b"))

			s = s.Next(StringSymbol("c"))
			expectSequenceMatches(t, s, StringSymbol("a"), StringSymbol("b"), StringSymbol("c"))

			s = s.Next(StringSymbol("d"))
			expectSequenceMatches(t, s, StringSymbol("b"), StringSymbol("c"), StringSymbol("d"))
		})
	})
}

package markov

import (
	"math/rand"
	"testing"
)

func TestGenerator(t *testing.T) {

	t.Run("Get()", func(t *testing.T) {

		t.Run("returns symbols from tables by rotating the sequence", func(t *testing.T) {
			ts := NewInMemoryTableStore()
			m := NewModel(ts)

			seq := EmptySequence(2)
			m.Add(seq, "a")

			seq = seq.WithNext("a")
			m.Add(seq, "b")

			seq = seq.WithNext("b")
			m.Add(seq, "c")

			gen := NewGenerator(m, 2, rand.New(rand.NewSource(12345)))

			s, err := gen.Get()

			if err != nil {
				t.Fatalf("Error generating: %v", err)
			}
			if expected, actual := "a", s; expected != actual {
				t.Errorf("Expected symbol '%s' but was '%s'", expected, actual)
			}

			s, err = gen.Get()

			if err != nil {
				t.Fatalf("Error generating: %v", err)
			}
			if expected, actual := "b", s; expected != actual {
				t.Errorf("Expected symbol '%s' but was '%s'", expected, actual)
			}

			s, err = gen.Get()

			if err != nil {
				t.Fatalf("Error generating: %v", err)
			}
			if expected, actual := "c", s; expected != actual {
				t.Errorf("Expected symbol '%s' but was '%s'", expected, actual)
			}
		})
	})
}

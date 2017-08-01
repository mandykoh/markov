package markov

import (
	"bytes"
)

type SequenceKey string

type Sequence struct {
	Symbols []Symbol
}

func (s Sequence) Key() SequenceKey {
	var keyBytes bytes.Buffer

	for _, symbol := range s.Symbols {
		var symbolKey string
		if symbol != nil {
			symbolKey = string(symbol.Key())
		} else {
			symbolKey = ""
		}

		if keyBytes.Len() > 0 {
			keyBytes.WriteString("|")
		}
		keyBytes.WriteString(string(symbolKey))
	}

	return SequenceKey(keyBytes.String())
}

func (s Sequence) Order() uint {
	return uint(len(s.Symbols))
}

func (s Sequence) WithNext(nextSym Symbol) Sequence {
	order := uint(len(s.Symbols))
	next := EmptySequence(order)
	copy(next.Symbols, s.Symbols[1:order])
	next.Symbols[order-1] = nextSym
	return next
}

func EmptySequence(order uint) Sequence {
	return Sequence{Symbols: make([]Symbol, order)}
}

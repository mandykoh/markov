package markov

import (
	"bytes"
	"strings"
)

const sequenceKeyEscape = '\\'
const sequenceKeySeparator = '|'

type SequenceKey string

type Sequence struct {
	Symbols []string
}

func (s Sequence) Key() SequenceKey {
	var keyBytes bytes.Buffer

	for _, symbol := range s.Symbols {
		if keyBytes.Len() > 0 {
			keyBytes.WriteRune(sequenceKeySeparator)
		}

		symbol = strings.Replace(symbol, string(sequenceKeyEscape), `\\`, -1)
		symbol = strings.Replace(symbol, string(sequenceKeySeparator), `\|`, -1)
		keyBytes.WriteString(symbol)
	}

	return SequenceKey(keyBytes.String())
}

func (s Sequence) Order() uint {
	return uint(len(s.Symbols))
}

func (s Sequence) WithNext(nextSym string) Sequence {
	order := uint(len(s.Symbols))
	next := EmptySequence(order)
	copy(next.Symbols, s.Symbols[1:order])
	next.Symbols[order-1] = nextSym
	return next
}

func EmptySequence(order uint) Sequence {
	return Sequence{Symbols: make([]string, order)}
}

func SequenceWith(symbols ...string) Sequence {
	return Sequence{Symbols: symbols}
}

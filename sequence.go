package markov

type Sequence struct {
	Symbols []Symbol
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

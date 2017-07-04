package markov

type Symbol interface {
	Equals(Symbol) bool
}

type StringSymbol string

func (s StringSymbol) Equals(sym Symbol) bool {
	switch v := sym.(type) {
	case StringSymbol:
		return v == s
	default:
		return false
	}
}

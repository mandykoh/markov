package markov

type StringSymbol string

func (s StringSymbol) Equals(sym Symbol) bool {
	switch v := sym.(type) {
	case StringSymbol:
		return v == s
	default:
		return false
	}
}

func (s StringSymbol) Key() SymbolKey {
	return SymbolKey(s)
}

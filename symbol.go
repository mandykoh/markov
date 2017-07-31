package markov

type SymbolKey string

type Symbol interface {
	Equals(Symbol) bool
	Key() SymbolKey
}

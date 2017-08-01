package markov

type Accumulator struct {
	Model           *Model
	CurrentSequence Sequence
}

func (acc *Accumulator) Add(symbol string) (err error) {
	err = acc.Model.Add(acc.CurrentSequence, symbol)
	if err == nil {
		acc.CurrentSequence = acc.CurrentSequence.WithNext(symbol)
	}
	return
}

func NewAccumulator(m *Model, order uint) *Accumulator {
	return &Accumulator{
		Model:           m,
		CurrentSequence: EmptySequence(order),
	}
}

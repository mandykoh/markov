package markov

type Accumulator struct {
	Model           *Model
	CurrentSequence Sequence
}

func (acc *Accumulator) Add(s Symbol) (err error) {
	err = acc.Model.Add(acc.CurrentSequence, s)
	if err == nil {
		acc.CurrentSequence = acc.CurrentSequence.WithNext(s)
	}
	return
}

func NewAccumulator(m *Model, order uint) *Accumulator {
	return &Accumulator{
		Model:           m,
		CurrentSequence: EmptySequence(order),
	}
}

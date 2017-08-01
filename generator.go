package markov

type SampleSource interface {
	Float64() float64
}

type Generator struct {
	Model           *Model
	CurrentSequence Sequence
	SampleSource    SampleSource
}

func (g *Generator) Get() (s Symbol, err error) {
	s, err = g.Model.Sample(g.CurrentSequence, g.SampleSource.Float64())
	if err == nil {
		g.CurrentSequence = g.CurrentSequence.WithNext(s)
	}
	return
}

func NewGenerator(m *Model, order uint, s SampleSource) *Generator {
	return &Generator{
		Model:           m,
		CurrentSequence: EmptySequence(order),
		SampleSource:    s,
	}
}

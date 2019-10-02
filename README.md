# markov

[![GoDoc](https://godoc.org/github.com/mandykoh/markov?status.svg)](https://godoc.org/github.com/mandykoh/markov)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/markov)](https://goreportcard.com/report/github.com/mandykoh/markov)
[![Build Status](https://travis-ci.org/mandykoh/markov.svg?branch=master)](https://travis-ci.org/mandykoh/markov)

Markov model library for Go

## Examples

Create a store for holding your Markov table data. Two built-in types are provided:

```go
store := markov.NewInMemoryTableStore()
```

```go
store, err := markov.NewBoltTableStore("path/to/store")
```

Stores may need to be closed after use:

```go
defer store.Close()
```

Create a Markov model using the store:

```go
model := markov.NewModel(store)
```

Use an Accumulator to build the model. The Accumulator needs the order of the model to build (here, we’re specifying a 3rd-order Markov model, meaning that each symbol is dependent on the three symbols that come before it):

```go
acc := markov.NewAccumulator(model, 3)
```

Symbols can be added in sequence to the accumulator. Symbols are strings, but can represent characters, words, events, or any abstract element in an ordered sequence.

```go
err := acc.Add("first symbol")
…
err = acc.Add("second symbol")
…
err = acc.Add("third symbol")
```

An empty string signifies the “end” of a sequence of symbols, if such a thing makes sense for your model:

```go
err := acc.Add("")
```

Use a Generator to sample from the model. The Generator needs to know the order of the model for sampling (usually the same order as used for the Accumulator), and also a `SampleSource` which determines how samples are drawn (here, we use a seeded pseudorandom number generator):

```go
gen := markov.NewGenerator(model, 3, rand.New(rand.NewSource(int64(12345))))
```

Sequences of symbols can be generated by sampling from the Generator:

```go
symbol, err := gen.Get()
```

If the end of a sequence is reached, the empty symbol ("") is returned with no error.

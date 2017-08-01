package markov

import "errors"

var ErrTableNotFound = errors.New("table not found")

type TableStore interface {
	Get(key SequenceKey, dest *Table) error
	Put(key SequenceKey, table *Table) error
}

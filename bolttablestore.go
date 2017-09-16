package markov

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
)

const BOLT_BUCKET = "markov"

type BoltTableStore struct {
	Bolt *bolt.DB
}

func (ts *BoltTableStore) Close() error {
	return ts.Bolt.Close()
}

func (ts *BoltTableStore) Get(k SequenceKey, dest *Table) error {
	key := ts.hashKey(string(k))
	return ts.Bolt.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(BOLT_BUCKET)).Get([]byte(string(key)))
		if len(data) > 0 {
			return json.Unmarshal(data, dest)
		}

		return ErrTableNotFound
	})
}

func (ts *BoltTableStore) Put(k SequenceKey, table *Table) error {
	key := ts.hashKey(string(k))
	return ts.Bolt.Update(func(tx *bolt.Tx) error {
		data, err := json.Marshal(table)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(BOLT_BUCKET)).Put([]byte(string(key)), data)
	})
}

func (ts *BoltTableStore) hashKey(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

func NewBoltTableStore(filePath string) (*BoltTableStore, error) {
	var db *bolt.DB
	var err error

	db, err = bolt.Open(filePath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BOLT_BUCKET))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &BoltTableStore{
		Bolt: db,
	}, nil
}

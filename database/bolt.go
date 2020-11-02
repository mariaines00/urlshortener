package database

import (
	"time"

	"github.com/boltdb/bolt"
)

// Storage is an interface which will be implmented by each storage
// e.g. bolt, sqlite
type Storage interface {
	GetEntryByID(string) (*Entry, error)
	DeleteEntry(string) error
	CreateEntry(Entry, string, string) error
	Close() error
}

// Entry is the data set which is stored in the DB as JSON
type Entry struct {
	URL         string
	RemoteAddr  string `json:",omitempty"`
	DeletionURL string `json:",omitempty"`
}

var (
	shortenedBucket = []byte("shortened")
)

// BoltStore implements the stores.Storage interface
type BoltStore struct {
	db *bolt.DB
}

// New returns a bolt store which implements the stores.Storage interface
func New(path string) (*BoltStore, error) {
	db, err := bolt.Open(path, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(shortenedBucket); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &BoltStore{
		db: db,
	}, nil
}

// Close closes the bolt database
func (b *BoltStore) Close() error {
	return b.db.Close()
}

// CreateEntry creates an entry by a given ID and returns an error
func (b *BoltStore) CreateEntry(entry Entry, id string) error {

}

// DeleteEntry deleted an entry by a given ID and returns an error
func (b *BoltStore) DeleteEntry(id string) error {

}

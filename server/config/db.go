package config

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"../shared"

	"github.com/boltdb/bolt"
)

var (
	shortenedBucket = []byte("shortened")
)

var DB *bolt.DB

// Init starts the database and created the server if needed
func Init(path string) error {
	DB, err := bolt.Open(path, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	err = DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(shortenedBucket); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}

	log.Println("Server connected to database")
	return nil
}

// Close closes the bolt database
func Close() error {
	return DB.Close()
}

// GetEntryByID returns a entry and an error by the shorted ID
func GetEntryByID(id string) (*shared.Entry, error) {
	raw := []byte{}
	err := DB.View(func(tx *bolt.Tx) error {
		raw = tx.Bucket(shortenedBucket).Get([]byte(id))
		if raw == nil {
			return errors.New("Nothing there")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var entry *shared.Entry
	return entry, json.Unmarshal(raw, &entry)
}

// CreateEntry creates an entry by a given ID and returns an error
func CreateEntry(id string, entry shared.Entry) error {
	entryRaw, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		if raw := bucket.Get([]byte(id)); raw != nil {
			return errors.New("entry already exists")
		}
		if err := bucket.Put([]byte(id), entryRaw); err != nil {
			return err
		}
		return nil
	})

	return err
}

// DeleteEntry deleted an entry by a given ID and returns an error
func DeleteEntry(id string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)

		if bucket.Get([]byte(id)) == nil {
			return errors.New("entry already deleted")
		}
		if err := bucket.Delete([]byte(id)); err != nil {
			return err
		}
		return nil
	})
	return err
}

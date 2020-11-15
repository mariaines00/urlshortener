package config

import (
	"encoding/json"
	"errors"
	"fmt"
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
	var err error
	DB, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	err = DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(shortenedBucket); err != nil {
			return err
		}
		log.Println("Created bucket", string(shortenedBucket))
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
		bucket := tx.Bucket(shortenedBucket)
		raw = bucket.Get([]byte(id))
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

	return DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		if raw := bucket.Get([]byte(id)); raw != nil {
			return errors.New("entry already exists")
		}
		if err := bucket.Put([]byte(id), entryRaw); err != nil {
			return err
		}
		return nil
	})
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

func GetAll() ([]shared.Entry, error) {
	a := []shared.Entry{}

	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(shortenedBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			e := shared.Entry{}
			_ = json.Unmarshal(v, &e)
			a = append(a, e)
		}
		return nil
	})

	fmt.Println(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

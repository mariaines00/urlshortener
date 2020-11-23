package config

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/mariaines00/urlshortener/shared"

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
		bucket, err := tx.CreateBucketIfNotExists(shortenedBucket)
		if err != nil {
			return err
		}
		err = bucket.SetSequence(1000)
		if err != nil {
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
	entry := shared.Entry{}
	raw := []byte{}

	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		raw = bucket.Get([]byte(id))
		if raw == nil {
			return shared.NewHTTPError(nil, 404, "Entry does not exist")
		}
		return nil
	})

	if err != nil {
		return &entry, err
	}

	return &entry, json.Unmarshal(raw, &entry)
}

// CreateEntry creates an entry by a given ID and returns an error
func CreateEntry(id string, entry shared.Entry) error {
	entryRaw, err := json.Marshal(entry)
	if err != nil {
		return shared.NewHTTPError(err, 500, "Something went wrong")
	}

	return DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		if raw := bucket.Get([]byte(id)); raw != nil {
			return shared.NewHTTPError(err, 400, "Entry Already Exists")
		}
		_, _ = bucket.NextSequence() //updates the counter but discard the value
		if err := bucket.Put([]byte(id), entryRaw); err != nil {
			return shared.NewHTTPError(err, 500, "Something went wrong")
		}
		return nil
	})
}

// DeleteEntry deleted an entry by a given ID and returns an error
func DeleteEntry(id string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)

		if bucket.Get([]byte(id)) == nil {
			return shared.NewHTTPError(nil, 404, "Entry does not exist")
		}
		if err := bucket.Delete([]byte(id)); err != nil {
			return shared.NewHTTPError(err, 500, "Something went wrong")
		}
		return nil
	})
	return err
}

func getAll() ([]shared.Entry, error) {
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

	if err != nil {
		return nil, err
	}
	return a, nil
}

// IncreaseHits increments the hits counter and sets the lat acessed time
func IncreaseHits(id string) error {
	entry, err := GetEntryByID(id)
	if err != nil {
		return err
	}
	entry.Hits++
	entry.LastAccess = time.Now().UTC()
	raw, err := json.Marshal(entry)
	if err != nil {
		return shared.NewHTTPError(err, 500, "Something went wrong")
	}
	return DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		if err := bucket.Put([]byte(id), raw); err != nil {
			return shared.NewHTTPError(err, 500, "Something went wrong")
		}
		return nil
	})
}

/* Helpers */

// GetSequence return the current auto-incremental id for the db
func GetSequence() (int, error) {
	n := 0
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(shortenedBucket)
		n64 := bucket.Sequence()
		n = int(n64)
		return nil
	})

	return n, shared.NewHTTPError(err, 500, "Something went wrong")
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

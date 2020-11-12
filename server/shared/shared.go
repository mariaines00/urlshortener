package shared

import "time"

// Entry is the data set which is stored in the DB as JSON
type Entry struct {
	Path                  string
	OutsideAddr           string
	Hits                  int       `json:",omitempty"`
	CreatedAt, LastAccess time.Time `json:",omitempty"`
}

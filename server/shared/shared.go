package shared

import "time"

// Entry is the data set which is stored in the DB as JSON
type Entry struct {
	Path                  string
	OutsideAddr           string
	Hits                  int
	CreatedAt, LastAccess time.Time
}

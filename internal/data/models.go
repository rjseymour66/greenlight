package data

import (
	"database/sql"
	"errors"
)

var (
	// ErrRecordNotFound when no record found.
	ErrRecordNotFound = errors.New("record not found")
)

// Models wriaps all models to hold all models as the application grows.
type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

// NewModels returns a Models object with db connections.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: &MovieModel{DB: db},
	}
}

// NewMockModels is a helper function which returns a Models instance
// containing the mock models only.
func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}

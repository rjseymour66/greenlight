package data

import (
	"database/sql"
	"errors"
)

var (
	// ErrRecordNotFound when no record found.
	ErrRecordNotFound = errors.New("record not found")
	// ErrEditConflict when more than one client tries to edit a db record at the same time.
	ErrEditConflict = errors.New("edit conflict")
)

// Models wriaps all models to hold all models as the application grows.
// create the interface if you need to create mocks for testing
// type Models struct {
// 	Movies interface {
// 		Insert(movie *Movie) error
// 		Get(id int64) (*Movie, error)
// 		Update(movie *Movie) error
// 		Delete(id int64) error
// 		GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
// 	}
// }

type Models struct {
	Movies MovieModel
	Tokens TokenModel
	Users  UserModel
}

// NewModels returns a Models object with db connections.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Tokens: TokenModel{DB: db},
		Users:  UserModel{DB: db},
	}
}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

/*
// NewMockModels is a helper function which returns a Models instance
// containing the mock models only.
func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
*/

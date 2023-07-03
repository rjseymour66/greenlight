package data

import (
	"database/sql"
	"greenlight/internal/validator"
	"time"

	"github.com/lib/pq"
)

// Movie represents a movie you can research.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// ValidateMovie makes sure the movie value is well-formed.
func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	// Note that we're using the Unique helper in the line below to check that all
	// values in the input.Genres slice are unique.
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

// MovieModel connects to the DB.
type MovieModel struct {
	DB *sql.DB
}

// Insert inserts the movie.
func (m *MovieModel) Insert(movie *Movie) error {
	query := `
	INSERT INTO movies (title, year, runtime, genres)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Get gets the movie.
func (m *MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Update updates the movie.
func (m *MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete deletes the movie.
func (m *MovieModel) Delete(id int64) error {
	return nil
}

// MockMovieModel for unit testing.
type MockMovieModel struct{}

// Insert inserts the movie.
func (m MockMovieModel) Insert(movie *Movie) error {
	// Mock the action...
	return nil
}

// Get gets the movie.
func (m MockMovieModel) Get(id int64) (*Movie, error) {
	// Mock the action...
	return nil, nil
}

// Update updates the movie.
func (m MockMovieModel) Update(movie *Movie) error {
	// Mock the action...
	return nil
}

// Delete deletes the movie.
func (m MockMovieModel) Delete(id int64) error {
	// Mock the action...
	return nil
}

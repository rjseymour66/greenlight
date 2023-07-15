package data

// Filters stores filters taken from query string parameters.
type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

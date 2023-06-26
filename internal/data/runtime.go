package data

import (
	"fmt"
	"strconv"
)

// Runtime is the length of a movie
type Runtime int32

// MarshalJSON is a custom implemenation that satisfies the json.Marshaler interface.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

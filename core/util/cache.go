package util

import (
	"encoding/json"
	"fmt"
)

// GenerateCacheKey generates a cache key based on the input parameters
func GenerateCacheKey(prefix string, params any) string {
	return fmt.Sprintf("%s:%v", prefix, params)
}

// Deserialize unmarshals the input data into the output interface
// Decode that JSON into the variable pointed to by result
func Deserialize(encodedJSON []byte, result any) error {
	return json.Unmarshal([]byte(encodedJSON), result)
}

// Serialize marshals the input data into an array of bytes
// Encoding input 'data' into JSON format.
func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

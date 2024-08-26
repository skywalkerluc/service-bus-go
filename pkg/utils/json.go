package utils

import (
	"encoding/json"
	"fmt"
)

// interface{} here represents an IO
func ToJSON(input interface{}) (string, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("error when serializing to Json: %w", err)
	}
	return string(b), nil
}

func FromJSON(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return fmt.Errorf("error when desserializing JSON: %w", err)
	}
	return nil
}

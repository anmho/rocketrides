package send

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteJSON[T any](w http.ResponseWriter, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func Read[T any](data io.ReadCloser) (T, error) {
	var v T
	if err := json.NewDecoder(data).Decode(&v); err != nil {
		return v, fmt.Errorf("decoding json: %w", err)
	}
	return v, nil
}

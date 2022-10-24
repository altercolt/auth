package web

import (
	"encoding/json"
	"io"
	"net/http"
)

func Decode(r *http.Request, val any) error {
	reader := io.LimitReader(r.Body, 5000000)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

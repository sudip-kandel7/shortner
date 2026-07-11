package utils

import (
	"net/http"
	"io/ioutils"
	"encoding/json"
)


func ParseBody(r *http.Request, x interface{}) error {
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body)
    if err != nil {
        return err
    }

    return json.Unmarshal(body, x)
}
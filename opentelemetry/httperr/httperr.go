package httperr

import (
	"net/http"

	"github.com/goccy/go-json"
)

type _error struct {
	Message string `json:"message"`
}

func NewError(w http.ResponseWriter, err error, code int) {
	parsed, err := json.Marshal(_error{Message: err.Error()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(code)
	_, err = w.Write(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

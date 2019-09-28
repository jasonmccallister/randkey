package api

import (
	"encoding/json"
	"net/http"

	"github.com/jasonmccallister/randkey/pkg/key"
)

type response struct {
	Key string `json:"key"`
}

// Handler is the endpoint that creates the serverless function for Zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Key: key.Generate(8),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

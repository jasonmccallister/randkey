package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jasonmccallister/randkey/pkg/key"
)

type response struct {
	Key string `json:"key"`
}

// Handler is the endpoint that creates the serverless function for Zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	l := getLength(r)

	resp := response{
		Key: key.Generate(l),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func getLength(r *http.Request) int {
	param := r.URL.Query().Get("length")

	if param != "" {
		i, err := strconv.Atoi(param)
		if err != nil {
			return 8
		}

		return i
	}

	return 8
}

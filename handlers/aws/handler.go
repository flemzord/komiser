package aws

import (
	"encoding/json"
	"net/http"

	. "github.com/flemzord/komiser/services/aws"
	. "github.com/flemzord/komiser/services/cache"
)

type AWSHandler struct {
	cache    Cache
	multiple bool
	aws      AWS
}

func NewAWSHandler(cache Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws:      AWS{},
	}
	return &awsHandler
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

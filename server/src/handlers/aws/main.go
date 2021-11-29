package aws

import (
	"encoding/json"
	"net/http"

	. "github.com/flemzord/komiser/src/services/aws"
)

type AWSHandler struct {
	multiple bool
	aws      AWS
}

func NewAWSHandler() *AWSHandler {
	awsHandler := AWSHandler{
		aws: AWS{},
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

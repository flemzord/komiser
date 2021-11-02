package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"net/http"
)

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	response, err := handler.aws.DescribeCostAndUsage(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
	} else {
		respondWithJSON(w, 200, response.Total)
	}
}

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	response, err := handler.aws.DescribeCostAndUsage(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
	} else {
		respondWithJSON(w, 200, response.History)
	}
}

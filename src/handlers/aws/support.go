package aws

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) SupportOpenTicketsHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.OpenSupportTickets(cfg)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "support:DescribeCases is missing")
	} else {
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) SupportTicketsInLastSixMonthsHandlers(w http.ResponseWriter, r *http.Request) {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.TicketsInLastSixMonthsTickets(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "support:DescribeCases is missing")
	} else {
		respondWithJSON(w, 200, response)
	}

}

func (handler *AWSHandler) DescribeServiceLimitsChecks(w http.ResponseWriter, r *http.Request) {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.DescribeServiceLimitsChecks(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "support:DescribeTrustedAdvisorChecks is missing")
	} else {
		respondWithJSON(w, 200, response)
	}
}

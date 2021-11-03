package aws

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) SupportOpenTicketsHandler(c *fiber.Ctx) error {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.OpenSupportTickets(cfg)
	if err != nil {
		return c.JSON(fiber.Map{"error": "support:DescribeCases is missing"})
	} else {
		return c.JSON(response)

	}
}

func (handler *AWSHandler) SupportTicketsInLastSixMonthsHandlers(c *fiber.Ctx) error {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.TicketsInLastSixMonthsTickets(cfg)
	if err != nil {
		return c.JSON(fiber.Map{"error": "support:DescribeCases is missing"})
	} else {
		return c.JSON(response)

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

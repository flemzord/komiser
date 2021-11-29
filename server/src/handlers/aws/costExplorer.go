package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/gofiber/fiber/v2"
)

func (handler *AWSHandler) CurrentCostHandler(c *fiber.Ctx) error {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.DescribeCostAndUsage(cfg)
	if err != nil {
		return c.JSON(fiber.Map{"error": "ce:GetCostAndUsage is missing"})
	} else {
		return c.JSON(response.Total)

	}
}

func (handler *AWSHandler) CostAndUsageHandler(c *fiber.Ctx) error {
	cfg, err := external.LoadDefaultAWSConfig()

	response, err := handler.aws.DescribeCostAndUsage(cfg)
	if err != nil {
		return c.JSON(fiber.Map{"error": "ce:GetCostAndUsage is missing"})
	} else {
		return c.JSON(response.History)

	}
}

package aws

import (
	"context"
	"sort"
	"strconv"
	"time"

	awsClient "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	. "github.com/flemzord/komiser/src/models/aws"
)

func (aws AWS) DescribeCostAndUsage(cfg awsClient.Config) (Bill, error) {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, -6, 0).Format("2006-01-02")
	end := currentTime.Format("2006-01-02")
	cfg.Region = "us-east-1"
	svc := costexplorer.New(cfg)
	req := svc.GetCostAndUsageRequest(&costexplorer.GetCostAndUsageInput{
		Metrics:     []string{"BlendedCost"},
		Granularity: costexplorer.GranularityMonthly,
		TimePeriod: &costexplorer.DateInterval{
			Start: &start,
			End:   &end,
		},
		//Filter: []costexplorer.Expression{
		//	Dimensions: []costexplorer.Expression{
		//        {
		//            Key:    awsClient.String("RECORD_TYPE"),
		//            Values: []string{"Usage"},
		//        },
		//    },
		//},
		GroupBy: []costexplorer.GroupDefinition{
			costexplorer.GroupDefinition{
				Key:  awsClient.String("SERVICE"),
				Type: costexplorer.GroupDefinitionTypeDimension,
			},
		},
	})
	result, err := req.Send(context.Background())
	if err != nil {
		return Bill{}, err
	}

	costs := make([]Cost, 0)
	for _, res := range result.ResultsByTime {
		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)

		unit := "USD"

		groups := make([]Group, 0)
		for _, group := range res.Groups {
			amount, _ := strconv.ParseFloat(*group.Metrics["BlendedCost"].Amount, 64)
			groups = append(groups, Group{
				Key:    group.Keys[0],
				Amount: amount,
			})
			unit = *group.Metrics["BlendedCost"].Unit
		}

		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Amount > groups[j].Amount
		})

		costs = append(costs, Cost{
			Start:  start,
			End:    end,
			Unit:   unit,
			Groups: groups,
		})
	}

	var currentBill float64
	for _, group := range costs[len(costs)-1].Groups {
		currentBill += group.Amount
	}

	return Bill{
		Total:   currentBill,
		History: costs,
	}, nil
}

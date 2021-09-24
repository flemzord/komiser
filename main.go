package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	. "github.com/flemzord/komiser/handlers/aws"
	. "github.com/flemzord/komiser/services/cache"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

const (
	DefaultPort     = 3000
	DefaultDuration = 30
)

//go:embed dashboard/dist
var distFiles embed.FS

func startServer(port int, cache Cache, dataset string, multiple bool) {
	cache.Connect()

	awsHandler := NewAWSHandler(cache, multiple)

	r := mux.NewRouter()
	r.HandleFunc("/aws/profiles", awsHandler.ConfigProfilesHandler)
	r.HandleFunc("/aws/iam/users", awsHandler.IAMUsersHandler)
	r.HandleFunc("/aws/iam/account", awsHandler.IAMUserHandler)
	r.HandleFunc("/aws/cost/current", awsHandler.CurrentCostHandler)
	r.HandleFunc("/aws/cost/history", awsHandler.CostAndUsageHandler)
	r.HandleFunc("/aws/resources/regions", awsHandler.UsedRegionsHandler)
	r.HandleFunc("/aws/cloudwatch/alarms", awsHandler.CloudWatchAlarmsHandler)
	r.HandleFunc("/aws/ec2/regions", awsHandler.EC2InstancesHandler)
	r.HandleFunc("/aws/lambda/functions", awsHandler.LambdaFunctionHandler)
	r.HandleFunc("/aws/lambda/invocations", awsHandler.GetLambdaInvocationMetrics)
	r.HandleFunc("/aws/s3/buckets", awsHandler.S3BucketsHandler)
	r.HandleFunc("/aws/s3/size", awsHandler.S3BucketsSizeHandler)
	r.HandleFunc("/aws/s3/objects", awsHandler.S3BucketsObjectsHandler)
	r.HandleFunc("/aws/ebs", awsHandler.EBSHandler)
	r.HandleFunc("/aws/rds/instances", awsHandler.RDSInstanceHandler)
	r.HandleFunc("/aws/dynamodb/tables", awsHandler.DynamoDBTableHandler)
	r.HandleFunc("/aws/elasticache/clusters", awsHandler.ElasticacheClustersHandler)
	r.HandleFunc("/aws/vpc", awsHandler.VPCHandler)
	r.HandleFunc("/aws/acl", awsHandler.ACLHandler)
	r.HandleFunc("/aws/route_tables", awsHandler.RouteTableHandler)
	r.HandleFunc("/aws/cloudfront/requests", awsHandler.CloudFrontRequestsHandler)
	r.HandleFunc("/aws/cloudfront/distributions", awsHandler.CloudFrontDistributionsHandler)
	r.HandleFunc("/aws/apigateway/requests", awsHandler.APIGatewayRequestsHandler)
	r.HandleFunc("/aws/apigateway/apis", awsHandler.APIGatewayRestAPIsHandler)
	r.HandleFunc("/aws/elb/requests", awsHandler.ELBRequestsHandler)
	r.HandleFunc("/aws/elb/family", awsHandler.ElasticLoadBalancerHandler)
	r.HandleFunc("/aws/kms", awsHandler.KMSKeysHandler)
	r.HandleFunc("/aws/key_pairs", awsHandler.KeyPairHandler)
	r.HandleFunc("/aws/security_groups", awsHandler.SecurityGroupHandler)
	r.HandleFunc("/aws/security_groups/unrestricted", awsHandler.ListUnrestrictedSecurityGroups)
	r.HandleFunc("/aws/acm/certificates", awsHandler.APIGatewayListCertificatesHandler)
	r.HandleFunc("/aws/acm/expired", awsHandler.APIGatewayExpiredCertificatesHandler)
	r.HandleFunc("/aws/sqs/messages", awsHandler.GetNumberOfMessagesSentAndDeletedSQSHandler)
	r.HandleFunc("/aws/sqs/queues", awsHandler.SQSQueuesHandler)
	r.HandleFunc("/aws/sns/topics", awsHandler.SNSTopicsHandler)
	r.HandleFunc("/aws/mq/brokers", awsHandler.ActiveMQBrokersHandler)
	r.HandleFunc("/aws/kinesis/streams", awsHandler.KinesisListStreamsHandler)
	r.HandleFunc("/aws/kinesis/shards", awsHandler.KinesisListShardsHandler)
	r.HandleFunc("/aws/glue/crawlers", awsHandler.GlueGetCrawlersHandler)
	r.HandleFunc("/aws/glue/jobs", awsHandler.GlueGetJobsHandler)
	r.HandleFunc("/aws/datapipeline/pipelines", awsHandler.DataPipelineListPipelines)
	r.HandleFunc("/aws/es/domains", awsHandler.ESListDomainsHandler)
	r.HandleFunc("/aws/swf/domains", awsHandler.SWFListDomainsHandler)
	r.HandleFunc("/aws/support/open", awsHandler.SupportOpenTicketsHandler)
	r.HandleFunc("/aws/support/history", awsHandler.SupportTicketsInLastSixMonthsHandlers)
	r.HandleFunc("/aws/ecs", awsHandler.ECSHandler)
	r.HandleFunc("/aws/route53/zones", awsHandler.Route53HostedZonesHandler)
	r.HandleFunc("/aws/route53/records", awsHandler.Route53ARecordsHandler)
	r.HandleFunc("/aws/logs/volume", awsHandler.LogsVolumeHandler)
	r.HandleFunc("/aws/cloudtrail/sign_in_event", awsHandler.CloudTrailConsoleSignInEventsHandler)
	r.HandleFunc("/aws/cloudtrail/source_ip", awsHandler.CloudTrailConsoleSignInSourceIpEventsHandler)
	r.HandleFunc("/aws/lambda/errors", awsHandler.GetLambdaErrorsMetrics)
	r.HandleFunc("/aws/ec2/scheduled", awsHandler.ScheduledEC2Instances)
	r.HandleFunc("/aws/ec2/reserved", awsHandler.ReservedEC2Instances)
	r.HandleFunc("/aws/ec2/spot", awsHandler.SpotEC2Instances)
	r.HandleFunc("/aws/cost/instance_type", awsHandler.CostAndUsagePerInstanceTypeHandler)
	r.HandleFunc("/aws/eks/clusters", awsHandler.EKSClustersHandler)
	r.HandleFunc("/aws/logs/retention", awsHandler.MaximumLogsRetentionPeriodHandler)
	r.HandleFunc("/aws/nat/traffic", awsHandler.GetNatGatewayTrafficHandler)
	r.HandleFunc("/aws/iam/organization", awsHandler.DescribeOrganizationHandler)
	r.HandleFunc("/aws/service/limits", awsHandler.DescribeServiceLimitsChecks)
	r.HandleFunc("/aws/s3/empty", awsHandler.GetEmptyBucketsHandler)
	r.HandleFunc("/aws/eip/detached", awsHandler.ElasticIPHandler)
	r.HandleFunc("/aws/redshift/clusters", awsHandler.DescribeRedshiftClustersHandler)
	r.HandleFunc("/aws/vpc/subnets", awsHandler.DescribeSubnetsHandler)
	r.HandleFunc("/aws/cost/forecast", awsHandler.DescribeForecastPriceHandler)

	sub, _ := fs.Sub(distFiles, "dashboard/dist")
	r.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	headersOk := handlers.AllowedHeaders([]string{"profile"})
	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(headersOk)(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Komiser"
	app.Version = "2.4.0"
	app.Usage = "Cloud Environment Inspector"
	app.Copyright = "Komiser - https://komiser.io"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Maxence Maireaux",
			Email: "maxence@maireaux.fr",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start server",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Usage: "Server port",
					Value: DefaultPort,
				},
				cli.IntFlag{
					Name:  "duration, d",
					Usage: "Cache expiration time",
					Value: DefaultDuration,
				},
				cli.StringFlag{
					Name:  "redis, r",
					Usage: "Redis server",
				},
				cli.StringFlag{
					Name:  "dataset, ds",
					Usage: "BigQuery Bill dataset",
				},
				cli.BoolFlag{
					Name:  "multiple, m",
					Usage: "Enable multiple AWS accounts",
				},
			},
			Action: func(c *cli.Context) error {
				port := c.Int("port")
				duration := c.Int("duration")
				redis := c.String("redis")
				dataset := c.String("dataset")
				multiple := c.Bool("multiple")

				var cache Cache

				if port == 0 {
					port = DefaultPort
				}
				if duration == 0 {
					duration = DefaultDuration
				}

				if redis == "" {
					cache = &Memory{
						Expiration: time.Duration(duration),
					}
				} else {
					cache = &Redis{
						Addr:       redis,
						Expiration: time.Duration(duration),
					}
				}

				startServer(port, cache, dataset, multiple)
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found %q !", command)
	}
	app.Run(os.Args)
}

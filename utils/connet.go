package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetConnection() dynamodb.DynamoDB {
	// snippet-start:[dynamodb.go.create_new_table.session]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		// TODO: connection
		//EC2IMDSEndpointMode: dynamodb.Endpoint{
		//	Address:              nil,
		//	CachePeriodInMinutes: nil,
		//},
	}))

	return *dynamodb.New(sess)
}

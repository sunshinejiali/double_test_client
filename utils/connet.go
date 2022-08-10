package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	ENDPOINT_AWS = "http://127.0.0.1:8000"
	ENDPOINT_TI  = "http://127.0.0.1:8080"
)

func GetLocalConnection() dynamodb.DynamoDB {
	// snippet-start:[dynamodb.go.create_new_table.session]
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(ENDPOINT_AWS),
		//Credentials: credentials.NewSharedCredentials("", "385595570414_DBaaS-DevUser-Role"),
		//Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, TOKEN),
	}))
	return *dynamodb.New(sess)
}

func GetTestConnection() dynamodb.DynamoDB {
	// snippet-start:[dynamodb.go.create_new_table.session]
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(ENDPOINT_TI),
		//Credentials: credentials.NewSharedCredentials("", "385595570414_DBaaS-DevUser-Role"),
		//Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, TOKEN),
	}))
	return *dynamodb.New(sess)
}

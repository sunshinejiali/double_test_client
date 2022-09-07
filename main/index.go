package main

import (
	"context"
	"double_test_client/judge"
	"double_test_client/log"
	"double_test_client/utils"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func query(router *gin.Engine) {
	fmt.Println("start test 'query':")
	// TODO: add value
	var (
		input *dynamodb.QueryInput
		//attributesToGet           []*string
		//conditionalOperator       *string
		//consistentRead            *bool
		//exclusiveStartKey         map[string]*dynamodb.AttributeValue
		//expressionAttributeNames  map[string]*string
		expressionAttributeValues map[string]*dynamodb.AttributeValue
		//filterExpression          *string
		//indexName              string
		keyConditionExpression string
		//keyConditions             map[string]*Condition
		//limit                     *int64
		//projectionExpression      *string
		//queryFilter               map[string]*Condition
		//returnConsumedCapacity    *string
		//scanIndexForward          *bool
		//selectAttribute string
		tableName string
		userId    string
		//topScore  string
		gameTitle string
	)

	expressionAttributeValues = make(map[string]*dynamodb.AttributeValue, 0)
	tableName = "Music"
	userId = "10"
	//topScore = "110001"
	gameTitle = "Happy Day"

	keyConditionExpression = "UserId = :userId and GameTitle = :gameTitle"
	//selectAttribute = "Select"
	expressionAttributeValues[":userId"] = &dynamodb.AttributeValue{
		S: &userId,
	}
	expressionAttributeValues[":gameTitle"] = &dynamodb.AttributeValue{
		S: &gameTitle,
	}
	//expressionAttributeValues[":topScore"] = &dynamodb.AttributeValue{
	//	N: &topScore,
	//}
	//indexName = "LocalSecondaryIndex"
	input = &dynamodb.QueryInput{
		//AttributesToGet:           nil,
		//ConditionalOperator:       nil,
		//ConsistentRead: &flag,
		//ExclusiveStartKey:         nil,
		ExpressionAttributeNames:  nil,
		ExpressionAttributeValues: expressionAttributeValues,
		//FilterExpression:          nil,
		//IndexName:              &indexName,
		KeyConditionExpression: &keyConditionExpression,
		//KeyConditions:             nil,
		//Limit:                     nil,
		//ProjectionExpression:      nil,
		//QueryFilter:               nil,
		//ReturnConsumedCapacity:    nil,
		//ScanIndexForward:          nil,
		//Select:    &selectAttribute,
		TableName: &tableName,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localQueryOutput, err := svc.Query(input)
	if err != nil {
		panic(err)
	}
	log.Info(context.TODO(), localQueryOutput)
	log.Info(context.TODO(), "The local query item test is finished")

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testQueryOutupt, err := sti.Query(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testQueryOutupt)
	log.Info(context.TODO(), "The server query item test is finished")

	judge.Query(*localQueryOutput, *testQueryOutupt)
}

func scan(router *gin.Engine) {
	fmt.Println("start test 'scan':")
	// TODO: add value
	var (
		input *dynamodb.ScanInput
		//attributesToGet           []*string
		//conditionalOperator       *string
		//consistentRead            *bool
		//exclusiveStartKey         map[string]*dynamodb.AttributeValue
		//expressionAttributeNames  map[string]*string
		//expressionAttributeValues map[string]*dynamodb.AttributeValue
		//filterExpression          *string
		//indexName                 *string
		//limit                     *int64
		//projectionExpression      *string
		//scanFilter                map[string]*Condition
		returnConsumedCapacity string
		//scanIndexForward          *bool
		//select                    *string
		tableName string
		//segment                   *int64
		//totalSegment              *int64
	)

	tableName = "Music"
	returnConsumedCapacity = "Total"
	input = &dynamodb.ScanInput{
		//AttributesToGet:           nil,
		//ConditionalOperator:       nil,
		//ConsistentRead:            nil,
		//ExclusiveStartKey:         nil,
		//ExpressionAttributeNames:  nil,
		//ExpressionAttributeValues: nil,
		//FilterExpression:          nil,
		//IndexName:                 nil,
		//Limit:                     nil,
		//ProjectionExpression:      nil,
		ReturnConsumedCapacity: &returnConsumedCapacity,
		//ScanFilter:                nil,
		//Segment:                   nil,
		//Select:                    nil,
		TableName: &tableName,
		//TotalSegments:             nil,
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testScanOutupt, err := sti.Scan(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testScanOutupt)
	log.Info(context.TODO(), "The server scan item test is finished")

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localScanOutput, err := svc.Scan(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), localScanOutput)
	log.Info(context.TODO(), "The local scan item test is finished")

	judge.Scan(*localScanOutput, *testScanOutupt)
}

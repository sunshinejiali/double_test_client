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

/*
	aws dynamodb query \
	    --table-name Music \
	    --key-condition-expression "Artist = :name" \
	    --expression-attribute-values  '{":name":{"S":"Acme Band"}}'  --endpoint-url http://localhost:8000

	{
	    "Items": [
	        {
	            "Artist": {
	                "S": "Acme Band"
	            },
	            "Awards": {
	                "N": "10"
	            },
	            "AlbumTitle": {
	                "S": "Updated Album Title"
	            },
	            "SongTitle": {
	                "S": "Happy Day"
	            }
	        }
	    ],
	    "Count": 1,
	    "ScannedCount": 1,
	    "ConsumedCapacity": null
	}
*/
func query(router *gin.Engine) {
	// region
	// countItem
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
		//indexName                 *string
		keyConditionExpression string
		//keyConditions             map[string]*Condition
		//limit                     *int64
		//projectionExpression      *string
		//queryFilter               map[string]*Condition
		//returnConsumedCapacity    *string
		//scanIndexForward          *bool
		//select                    *string
		tableName string
		artist    string
	)

	expressionAttributeValues = make(map[string]*dynamodb.AttributeValue, 0)
	// add the value
	//flag := true
	tableName = "Music"
	artist = "Acme Band"
	keyConditionExpression = "Artist = :name"
	expressionAttributeValues[":name"] = &dynamodb.AttributeValue{
		S: &artist,
	}
	input = &dynamodb.QueryInput{
		//AttributesToGet:           nil,
		//ConditionalOperator:       nil,
		//ConsistentRead: &flag,
		//ExclusiveStartKey:         nil,
		//ExpressionAttributeNames:  nil,
		ExpressionAttributeValues: expressionAttributeValues,
		//FilterExpression:          nil,
		//IndexName:                 nil,
		KeyConditionExpression: &keyConditionExpression,
		//KeyConditions:             nil,
		//Limit:                     nil,
		//ProjectionExpression:      nil,
		//QueryFilter:               nil,
		//ReturnConsumedCapacity:    nil,
		//ScanIndexForward:          nil,
		//Select:                    nil,
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

/*
aws dynamodb scan --table-name Music --return-consumed-capacity Total  --endpoint-url http://localhost:8000

	{
	    "Items": [
	        {
	            "Artist": {
	                "S": "Acme Band"
	            },
	            "Awards": {
	                "N": "10"
	            },
	            "AlbumTitle": {
	                "S": "Updated Album Title"
	            },
	            "SongTitle": {
	                "S": "Happy Day"
	            }
	        }
	    ],
	    "Count": 1,
	    "ScannedCount": 1,
	    "ConsumedCapacity": {
	        "TableName": "Music",
	        "CapacityUnits": 0.5
	    }
	}
*/
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

	// add the value
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

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localScanOutput, err := svc.Scan(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), localScanOutput)
	log.Info(context.TODO(), "The local scan item test is finished")

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testScanOutupt, err := sti.Scan(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testScanOutupt)
	log.Info(context.TODO(), "The server scan item test is finished")

	judge.Scan(*localScanOutput, *testScanOutupt)
}

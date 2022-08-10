package main

import (
	"double_test_client/judge"
	"double_test_client/utils"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

/*
func batchGetItem(router *gin.Engine) {
	fmt.Println("start test 'batchGetItem':")
	var (
		input                  *dynamodb.BatchGetItemInput
		requestItems           map[string]*dynamodb.KeysAndAttributes
		returnConsumedCapacity string
	)
	// add the value
	requestItems["Music"] = &dynamodb.KeysAndAttributes{
		AttributesToGet:         // nil,
		ConsistentRead:           nil,
		ExpressionAttributeNames: nil,
		Keys:                     nil,
		ProjectionExpression:     nil,
	}
	// TODO: add all values
	input = &dynamodb.BatchGetItemInput{
		RequestItems:           requestItems,
		ReturnConsumedCapacity: &returnConsumedCapacity,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localGetItemOutput, err := svc.BatchGetItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testGetItemOutput, err := sti.BatchGetItem(input)
	if err != nil {
		panic(err)
	}
	judge.BatchGetItem(*localGetItemOutput, *testGetItemOutput)
}
*/

/*
func batchWriteItem(router *gin.Engine) {
	fmt.Println("start test 'batchPutItem':")
	var (
		input                  *dynamodb.BatchWriteItemInput
		requestItems           map[string]*dynamodb.KeysAndAttributes
		returnConsumedCapacity string
		returnItemCollectionMetrics string
	)
	// add the value
	requestItems["Music"] = &dynamodb.KeysAndAttributes{
		AttributesToGet:         // nil,
		ConsistentRead:           nil,
		ExpressionAttributeNames: nil,
		Keys:                     nil,
		ProjectionExpression:     nil,
	}
	// TODO: add all values
	input = &dynamodb.BatchGetItemInput{
		RequestItems:           requestItems,
		ReturnConsumedCapacity: &returnConsumedCapacity,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localGetItemOutput, err := svc.BatchGetItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testGetItemOutput, err := sti.BatchGetItem(input)
	if err != nil {
		panic(err)
	}
	judge.BatchGetItem(*localGetItemOutput, *testGetItemOutput)
}
*/

func getItem(router *gin.Engine) {
	fmt.Println("start test 'getItem':")
	var (
		input *dynamodb.GetItemInput
		//attributesToGet          []*string
		//consistentRead           bool
		//expressionAttributeNames map[string]*string
		key map[string]*dynamodb.AttributeValue
		//projectionExpression     string
		//returnConsumedCapacity   string
		tableName string
		artist    string
		songTitle string
	)
	// add the value
	// TODO: add all values
	tableName = "Music"
	artist = "Acme Band"
	songTitle = "Happy Day"
	key["Artist"] = &dynamodb.AttributeValue{
		S: &artist,
	}
	key["SongTitle"] = &dynamodb.AttributeValue{
		S: &songTitle,
	}
	input = &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localGetItemOutput, err := svc.GetItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testGetItemOutput, err := sti.GetItem(input)
	if err != nil {
		panic(err)
	}
	judge.GetItem(*localGetItemOutput, *testGetItemOutput)
}

func putItem(router *gin.Engine) {
	fmt.Println("start test 'putItem':")
	var (
		input *dynamodb.PutItemInput
		//ConditionExpression         string
		//ConditionalOperator         string
		//Expected                    map[string]*ExpectedAttributeValue
		//ExpressionAttributeNames    map[string]*string
		//ExpressionAttributeValues   map[string]*AttributeValue
		item map[string]*dynamodb.AttributeValue
		//returnConsumedCapacity      string
		//returnItemCollectionMetrics string
		//returnValues                string
		tableName  string
		artist     string
		songTitle  string
		albumTitle string
		awards     string
	)
	// add the value
	// TODO: add all values
	tableName = "Music"
	artist = "Acme Band"
	songTitle = "Happy Day"
	albumTitle = "Songs About Life"
	awards = "10"
	item["Artist"] = &dynamodb.AttributeValue{
		S: &artist,
	}
	item["SongTitle"] = &dynamodb.AttributeValue{
		S: &songTitle,
	}
	item["AlbumTitle"] = &dynamodb.AttributeValue{
		S: &albumTitle,
	}
	item["Awards"] = &dynamodb.AttributeValue{
		N: &awards,
	}
	input = &dynamodb.PutItemInput{
		//ConditionExpression:         nil,
		//ConditionalOperator:         nil,
		//Expected:                    nil,
		//ExpressionAttributeNames:    nil,
		//ExpressionAttributeValues:   nil,
		Item: item,
		//ReturnConsumedCapacity:      &returnConsumedCapacity,
		//ReturnItemCollectionMetrics: &returnItemCollectionMetrics,
		//ReturnValues:                &returnValues,
		TableName: &tableName,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localPutItemOutput, err := svc.PutItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testPutItemOutput, err := sti.PutItem(input)
	if err != nil {
		panic(err)
	}
	judge.PutItem(*localPutItemOutput, *testPutItemOutput)
}

func updateItem(router *gin.Engine) {
	fmt.Println("start test 'updateItem':")
	var (
		input *dynamodb.UpdateItemInput
		//attributeUpdates            map[string]*AttributeValueUpdate
		//conditionExpression         *string
		//conditionalOperator         *string
		//expected                    map[string]*ExpectedAttributeValue
		//expressionAttributeNames    map[string]*string
		expressionAttributeValues map[string]*dynamodb.AttributeValue
		key                       map[string]*dynamodb.AttributeValue
		//returnConsumedCapacity      *string
		//returnItemCollectionMetrics *string
		returnValues     string
		tableName        string
		updateExpression string
		artist           string
		songTitle        string
		albumTitle       string
	)
	// add the value
	// TODO: add all values
	tableName = "Music"
	artist = "Acme Band"
	songTitle = "Happy Day"
	key["Artist"] = &dynamodb.AttributeValue{
		S: &artist,
	}
	key["SongTitle"] = &dynamodb.AttributeValue{
		S: &songTitle,
	}
	returnValues = "ALL_NEW"
	albumTitle = "Updated Album Title"
	expressionAttributeValues[":newval"] = &dynamodb.AttributeValue{
		S: &albumTitle,
	}
	updateExpression = "SET AlbumTitle = :newval"
	input = &dynamodb.UpdateItemInput{
		//AttributeUpdates:            nil,
		//ConditionExpression:         nil,
		//ConditionalOperator:         nil,
		//Expected:                    nil,
		//ExpressionAttributeNames:    nil,
		ExpressionAttributeValues: expressionAttributeValues,
		Key:                       key,
		//ReturnConsumedCapacity:      nil,
		//ReturnItemCollectionMetrics: nil,
		ReturnValues:     &returnValues,
		TableName:        &tableName,
		UpdateExpression: &updateExpression,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localUpdateItemOutput, err := svc.UpdateItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testUpdateItemOutput, err := sti.UpdateItem(input)
	if err != nil {
		panic(err)
	}
	judge.UpdateItem(*localUpdateItemOutput, *testUpdateItemOutput)
}

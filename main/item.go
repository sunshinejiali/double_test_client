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
	// TODO: add all values
	key = make(map[string]*dynamodb.AttributeValue, 0)
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

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testGetItemOutput, err := sti.GetItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testGetItemOutput)
	log.Info(context.TODO(), "The server get item test is finished")

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localGetItemOutput, err := svc.GetItem(input)
	if err != nil {
		panic(err)
	}
	log.Info(context.TODO(), localGetItemOutput)
	log.Info(context.TODO(), "The local get item test is finished")

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
	item = make(map[string]*dynamodb.AttributeValue, 0)

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

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testPutItemOutput, err := sti.PutItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testPutItemOutput)
	log.Info(context.TODO(), "The server put item test is finished")

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localPutItemOutput, err := svc.PutItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), localPutItemOutput)
	log.Info(context.TODO(), "The local put item test is finished")

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

	// TODO: add all values
	key = make(map[string]*dynamodb.AttributeValue, 0)
	expressionAttributeValues = make(map[string]*dynamodb.AttributeValue, 0)
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
	albumTitle = "Updated Album Titles"
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

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testUpdateItemOutput, err := sti.UpdateItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testUpdateItemOutput)
	log.Info(context.TODO(), "The server update item test is finished")

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localUpdateItemOutput, err := svc.UpdateItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), localUpdateItemOutput)
	log.Info(context.TODO(), "The local update item test is finished")

	judge.UpdateItem(*localUpdateItemOutput, *testUpdateItemOutput)
}

func deleteItem(router *gin.Engine) {
	fmt.Println("start test 'deleteItem':")
	var (
		input                        *dynamodb.DeleteItemInput
		key                          map[string]*dynamodb.AttributeValue
		tableName, artist, songTitle string
	)
	key = make(map[string]*dynamodb.AttributeValue, 0)

	// TODO: add all values
	tableName = "Music"
	artist = "Acme Band"
	songTitle = "Happy Day"
	key["SongTitle"] = &dynamodb.AttributeValue{
		S: &songTitle,
	}
	key["Artist"] = &dynamodb.AttributeValue{
		S: &artist,
	}
	input = &dynamodb.DeleteItemInput{
		//ConditionExpression:         nil,
		//ConditionalOperator:         nil,
		//Expected:                    nil,
		//ExpressionAttributeNames:    nil,
		//ExpressionAttributeValues:   nil,
		Key: key,
		//ReturnConsumedCapacity:      nil,
		//ReturnItemCollectionMetrics: nil,
		//ReturnValues:                nil,
		TableName: &tableName,
	}

	// ===== test dynamodb on tikv =====
	sti := utils.GetTestConnection()
	testDeleteItemOutput, err := sti.DeleteItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), testDeleteItemOutput)
	log.Info(context.TODO(), "The server put item test is finished")

	//  ===== test local dynamodb =====
	svc := utils.GetLocalConnection()
	localDeleteItemOutput, err := svc.DeleteItem(input)
	if err != nil {
		panic(err)
	}

	log.Info(context.TODO(), localDeleteItemOutput)
	log.Info(context.TODO(), "The local put item test is finished")

	judge.DeleteItem(*localDeleteItemOutput, *testDeleteItemOutput)
}

package main

import (
	"double_test_client/judge"
	"double_test_client/utils"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func batchGetItem() {

}

func batchWriteItem() {

}

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
	svc := utils.GetConnection()
	localGetItemOutput, err := svc.GetItem(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====

	router.GET("/getItem", func(context *gin.Context) {
		marInput, _ := json.Marshal(input)
		var body = strings.NewReader(string(marInput))
		response, err := http.Post("https://localhost:8989", "application/x-amz-json-1.0", body)
		if err != nil || response.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			panic(err)
		}
		reader := response.Body
		// TODO: extral attributes in header.
		//contentLength := response.ContentLength
		//contentType := response.Header.Get("Content-Type")
		//extraHeaders := map[string]string{
		//	//
		//}
		//context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		judge.GetItem(*localGetItemOutput, utils.IoReaderChangeBytes(reader))
	})
}

func putItem() {

}

func updateItem() {

}

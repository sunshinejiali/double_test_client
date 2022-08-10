package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"double_test_client/judge"
	"double_test_client/utils"
)

/*
	aws dynamodb create-table \
	    --table-name Music \
	    --attribute-definitions \
	        AttributeName=Artist,AttributeType=S \
	        AttributeName=SongTitle,AttributeType=S \
	    --key-schema \
	        AttributeName=Artist,KeyType=HASH \
	        AttributeName=SongTitle,KeyType=RANGE \
	    --provisioned-throughput \
	        ReadCapacityUnits=10,WriteCapacityUnits=5 \
	    --endpoint-url http://localhost:8000

	{
	    "TableDescription": {
	        "AttributeDefinitions": [
	            {
	                "AttributeName": "Artist",
	                "AttributeType": "S"
	            },
	            {
	                "AttributeName": "SongTitle",
	                "AttributeType": "S"
	            }
	        ],
	        "TableName": "Music",
	        "KeySchema": [
	            {
	                "AttributeName": "Artist",
	                "KeyType": "HASH"
	            },
	            {
	                "AttributeName": "SongTitle",
	                "KeyType": "RANGE"
	            }
	        ],
	        "TableStatus": "ACTIVE",
	        "CreationDateTime": "2022-08-10T10:06:50.095000+08:00",
	        "ProvisionedThroughput": {
	            "LastIncreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "LastDecreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "NumberOfDecreasesToday": 0,
	            "ReadCapacityUnits": 10,
	            "WriteCapacityUnits": 5
	        },
	        "TableSizeBytes": 0,
	        "ItemCount": 0,
	        "TableArn": "arn:aws:dynamodb:ddblocal:000000000000:table/Music"
	    }
	}
*/
func createtable(router *gin.Engine) {
	fmt.Println("start test 'createTable':")
	var (
		input                *dynamodb.CreateTableInput
		attributeDefinitions []*dynamodb.AttributeDefinition
		//billingMode           string
		keySchema             []*dynamodb.KeySchemaElement
		provisionedThroughput dynamodb.ProvisionedThroughput
		//tableClass            string
		//globalSecondaryIndex
		//localSecondaryIndex
		tableName string
	)

	// add the value
	tableName = "Music"
	attributeDefinitions = []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("SongTitle"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("Artist"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("AlbumTitle"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("Price"),
			AttributeType: aws.String("N"),
		},
		{
			AttributeName: aws.String("Genre"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("Year"),
			AttributeType: aws.String("N"),
		},
	}
	keySchema = []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("SongTitle"),
			KeyType:       aws.String("HASH"),
		},
		{
			AttributeName: aws.String("Artist"),
			KeyType:       aws.String("RANGE"),
		},
	}
	provisionedThroughput = dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	}
	// TODO: other optional attributes

	//billingMode = "PAY_PER_REQUEST"
	//globalSecondaryIndex := []*dynamodb.GlobalSecondaryIndex{
	//	{
	//		IndexName: aws.String("GlobalSecondaryIndex"),
	//		KeySchema: []*dynamodb.KeySchemaElement{
	//			{
	//				AttributeName: aws.String("Genre"),
	//				KeyType:       aws.String("S"),
	//			},
	//			{
	//				AttributeName: aws.String("SongTitle"),
	//				KeyType:       aws.String("S"),
	//			},
	//		},
	//	},
	//}
	//localSecondaryIndex := []*dynamodb.LocalSecondaryIndex{
	//	{
	//		IndexName: aws.String("LocalSecondaryIndex"),
	//		KeySchema: []*dynamodb.KeySchemaElement{
	//			{
	//				AttributeName: aws.String("AlbumTitle"),
	//				KeyType:       aws.String("S"),
	//			},
	//			{
	//				AttributeName: aws.String("SongTitle"),
	//				KeyType:       aws.String("S"),
	//			},
	//		},
	//	},
	//}
	//tableClass = aws.String("STANDARD")
	input = &dynamodb.CreateTableInput{
		AttributeDefinitions: attributeDefinitions,
		//BillingMode:            billingMode,
		//GlobalSecondaryIndexes: globalSecondaryIndex,
		KeySchema: keySchema,
		//LocalSecondaryIndexes:  localSecondaryIndex,
		ProvisionedThroughput: &provisionedThroughput,
		//TableClass:             tableClass,
		TableName: &tableName,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetConnection()
	localCreateTableOutput, err := svc.CreateTable(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	router.GET("/createTables", func(context *gin.Context) {
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
		//
		//}
		//context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		judge.CreateTable(*localCreateTableOutput, utils.IoReaderChangeBytes(reader))
	})
}

/*
aws dynamodb delete-table --table-name Music --endpoint-url http://localhost:8000

	{
	    "TableDescription": {
	        "AttributeDefinitions": [
	            {
	                "AttributeName": "Artist",
	                "AttributeType": "S"
	            },
	            {
	                "AttributeName": "SongTitle",
	                "AttributeType": "S"
	            }
	        ],
	        "TableName": "Music",
	        "KeySchema": [
	            {
	                "AttributeName": "Artist",
	                "KeyType": "HASH"
	            },
	            {
	                "AttributeName": "SongTitle",
	                "KeyType": "RANGE"
	            }
	        ],
	        "TableStatus": "ACTIVE",
	        "CreationDateTime": "2022-08-10T15:40:23.971000+08:00",
	        "ProvisionedThroughput": {
	            "LastIncreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "LastDecreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "NumberOfDecreasesToday": 0,
	            "ReadCapacityUnits": 10,
	            "WriteCapacityUnits": 5
	        },
	        "TableSizeBytes": 0,
	        "ItemCount": 0,
	        "TableArn": "arn:aws:dynamodb:ddblocal:000000000000:table/Music"
	    }
	}
*/
func deleteTable(router *gin.Engine) {
	fmt.Println("start test 'deleteTable':")
	var (
		input     *dynamodb.DeleteTableInput
		tableName string
	)

	// add the value
	tableName = "Music"
	input = &dynamodb.DeleteTableInput{
		TableName: &tableName,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetConnection()
	localDeleteTableOutput, err := svc.DeleteTable(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	router.GET("/deleteTables", func(context *gin.Context) {
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
		//
		//}
		//context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		judge.DeleteTable(*localDeleteTableOutput, utils.IoReaderChangeBytes(reader))
	})
}

/*
aws dynamodb update-table --table-name Person --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

	{
	    "TableDescription": {
	        "AttributeDefinitions": [
	            {
	                "AttributeName": "Country",
	                "AttributeType": "S"
	            },
	            {
	                "AttributeName": "Name",
	                "AttributeType": "S"
	            }
	        ],
	        "TableName": "Person",
	        "KeySchema": [
	            {
	                "AttributeName": "Country",
	                "KeyType": "HASH"
	            },
	            {
	                "AttributeName": "Name",
	                "KeyType": "RANGE"
	            }
	        ],
	        "TableStatus": "ACTIVE",
	        "CreationDateTime": "2022-08-10T15:40:50.794000+08:00",
	        "ProvisionedThroughput": {
	            "LastIncreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "LastDecreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "NumberOfDecreasesToday": 0,
	            "ReadCapacityUnits": 5,
	            "WriteCapacityUnits": 5
	        },
	        "TableSizeBytes": 0,
	        "ItemCount": 0,
	        "TableArn": "arn:aws:dynamodb:ddblocal:000000000000:table/Person"
	    }
	}
*/
func updateTable(router *gin.Engine) {
	fmt.Println("start test 'updateTable':")
	var (
		input                 *dynamodb.UpdateTableInput
		tableName             string
		provisionedThroughput dynamodb.ProvisionedThroughput
		//attributeDefinitions []*dynamodb.AttributeDefinition
		//billingMode           string
		//tableClass            string
		//globalSecondaryIndexUpdates []*GlobalSecondaryIndexUpdate
		//ReplicaUpdates       []*ReplicationGroupUpdate
	)

	// add the value
	tableName = "Music"
	provisionedThroughput = dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(5),
		WriteCapacityUnits: aws.Int64(5),
	}
	// TODO: update all values
	input = &dynamodb.UpdateTableInput{
		TableName:             &tableName,
		ProvisionedThroughput: &provisionedThroughput,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetConnection()
	localUpdateTableOutput, err := svc.UpdateTable(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	router.GET("/updateTable", func(context *gin.Context) {
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
		//
		//}
		//context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		judge.UpdateTable(*localUpdateTableOutput, utils.IoReaderChangeBytes(reader))
	})
}

/*
aws dynamodb describe-table --table-name Person --endpoint-url http://localhost:8000

	{
	    "Table": {
	        "AttributeDefinitions": [
	            {
	                "AttributeName": "Country",
	                "AttributeType": "S"
	            },
	            {
	                "AttributeName": "Name",
	                "AttributeType": "S"
	            }
	        ],
	        "TableName": "Person",
	        "KeySchema": [
	            {
	                "AttributeName": "Country",
	                "KeyType": "HASH"
	            },
	            {
	                "AttributeName": "Name",
	                "KeyType": "RANGE"
	            }
	        ],
	        "TableStatus": "ACTIVE",
	        "CreationDateTime": "2022-08-10T15:40:50.794000+08:00",
	        "ProvisionedThroughput": {
	            "LastIncreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "LastDecreaseDateTime": "1970-01-01T08:00:00+08:00",
	            "NumberOfDecreasesToday": 0,
	            "ReadCapacityUnits": 10,
	            "WriteCapacityUnits": 5
	        },
	        "TableSizeBytes": 0,
	        "ItemCount": 0,
	        "TableArn": "arn:aws:dynamodb:ddblocal:000000000000:table/Person"
	    }
	}
*/
func describeTable(router *gin.Engine) {
	fmt.Println("start test 'describeTable':")
	var (
		input     *dynamodb.DescribeTableInput
		tableName string
	)

	// add the value
	tableName = "Music"
	input = &dynamodb.DescribeTableInput{
		TableName: &tableName,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetConnection()
	localDescribeTableOutput, err := svc.DescribeTable(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	router.GET("/describeTable", func(context *gin.Context) {
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
		//
		//}
		//context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		judge.DescribeTable(*localDescribeTableOutput, utils.IoReaderChangeBytes(reader))
	})
}

/*
listTables(current data):

	{
		"Music":
		"Person":
		"Shopping":
	}

	aws dynamodb list-tables \
	 --exclusive-start-table-name Person \
	 --limit 1 \
	 --endpoint-url http://localhost:8000

	{
	    "TableNames": [
	        "Shopping"
	    ]
	}
*/
func listTables(router *gin.Engine) {
	fmt.Println("start test 'listTables':")
	var (
		input                   *dynamodb.ListTablesInput
		exclusiveStartTableName string
		limit                   int64
	)
	// add the value
	exclusiveStartTableName = "Music"
	limit = 1
	input = &dynamodb.ListTablesInput{
		ExclusiveStartTableName: &exclusiveStartTableName,
		Limit:                   &limit,
	}

	//  ===== test local dynamodb =====
	svc := utils.GetConnection()
	localListTablesOutput, err := svc.ListTables(input)
	if err != nil {
		panic(err)
	}

	// ===== test dynamodb on tikv =====
	router.GET("/listTables", func(context *gin.Context) {
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

		judge.ListTables(*localListTablesOutput, utils.IoReaderChangeBytes(reader))
	})
}

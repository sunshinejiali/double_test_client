package judge

import (
	"context"
	"double_test_client/log"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetItem(localGetItemOutput, testGetItemOutput dynamodb.GetItemOutput) {
	// judge
	if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
		panic("judgeGetItem test is fail: ConsumedCapacity is different.")
	}
	if len(localGetItemOutput.Item) != len(testGetItemOutput.Item) {
		panic("The size of Item is different.")
	}
	for index, name := range localGetItemOutput.Item {
		if name != testGetItemOutput.Item[index] {
			panic("The Item is different.")
		}
	}
	// result
	log.Info(context.TODO(), "getItem is successful!")
}

func PutItem(localGetItemOutput, testGetItemOutput dynamodb.PutItemOutput) {
	// judge
	if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
		panic("judgePutItem test is fail: ConsumedCapacity is different.")
	}
	if localGetItemOutput.ItemCollectionMetrics != testGetItemOutput.ItemCollectionMetrics {
		panic("judgePutItem test is fail: ItemCollectionMetrics is different.")
	}
	if len(localGetItemOutput.Attributes) != len(testGetItemOutput.Attributes) {
		panic("The size of Attributes is different.")
	}
	for index, name := range localGetItemOutput.Attributes {
		if name != testGetItemOutput.Attributes[index] {
			panic("The Attributes is different.")
		}
	}
	// result
	log.Info(context.TODO(), "putItem is successful!")
}

func UpdateItem(localGetItemOutput, testGetItemOutput dynamodb.UpdateItemOutput) {
	// judge
	if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
		panic("judgeUpdateItem test is fail: ConsumedCapacity is different.")
	}
	if localGetItemOutput.ItemCollectionMetrics != testGetItemOutput.ItemCollectionMetrics {
		panic("judgeUpdateItem test is fail: ItemCollectionMetrics is different.")
	}
	if len(localGetItemOutput.Attributes) != len(testGetItemOutput.Attributes) {
		panic("The size of Attributes is different.")
	}
	for index, name := range localGetItemOutput.Attributes {
		if name != testGetItemOutput.Attributes[index] {
			panic("The Attributes is different.")
		}
	}
	// result
	log.Info(context.TODO(), "updateItem is successful!")
}

/*
func BatchGetItem(localGetItemOutput, testGetItemOutput dynamodb.BatchGetItemOutput) {
	// judge
	if len(localGetItemOutput.ConsumedCapacity) != len(testGetItemOutput.ConsumedCapacity) {
		panic("The size of ConsumedCapacity is different.")
	}
	for index, capacity := range localGetItemOutput.ConsumedCapacity {
		if capacity != testGetItemOutput.ConsumedCapacity[index] {
			panic("The ConsumedCapacity is different.")
		}
	}
	if len(localGetItemOutput.Responses) != len(testGetItemOutput.Responses) {
		panic("The size of Responses is different.")
	}
	for index, response := range localGetItemOutput.Responses {
		testResponse := testGetItemOutput.Responses[index]
		if len(response) != len(testResponse) {
			panic("The size of response is different.")
		}
		for i, r := range response {
			for key, value := range r {
				temp := testResponse[i][key]
				if value != temp {
					panic("The AttributeValue is different.")
				}
			}
		}

	}
	if len(localGetItemOutput.UnprocessedKeys) != len(testGetItemOutput.UnprocessedKeys) {
		panic("The size of UnprocessedKeys is different.")
	}
	for index, key := range localGetItemOutput.UnprocessedKeys {
		if key != testGetItemOutput.UnprocessedKeys[index] {
			panic("The UnprocessedKeys is different.")
		}
	}
	// result
	log.Info(context.TODO(), "batchGetItem is successful!")
}
*/

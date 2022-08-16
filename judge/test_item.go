package judge

import (
	"context"
	"double_test_client/log"
	"double_test_client/utils"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sort"
)

func GetItem(localGetItemOutput, testGetItemOutput dynamodb.GetItemOutput) {
	// TODO:  judge
	//if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
	//	panic("judgeGetItem test is fail: ConsumedCapacity is different.")
	//}
	if len(localGetItemOutput.Item) != len(testGetItemOutput.Item) {
		panic("The size of Attributes is different.")
	}
	sortItem := make([]string, 0)
	for k := range localGetItemOutput.Item {
		sortItem = append(sortItem, k)
	}
	sort.Strings(sortItem)
	for _, k := range sortItem {
		if utils.Change(localGetItemOutput.Item[k]) != utils.Change(testGetItemOutput.Item[k]) {
			panic("The Attributes is different.")
		}
	}
	log.Info(context.TODO(), "getItem is successful!")
}

func PutItem(localGetItemOutput, testGetItemOutput dynamodb.PutItemOutput) {
	// TODO:  judge
	//if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
	//	panic("judgePutItem test is fail: ConsumedCapacity is different.")
	//}
	// TODO: judge
	//if localGetItemOutput.ItemCollectionMetrics != testGetItemOutput.ItemCollectionMetrics {
	//	panic("judgePutItem test is fail: ItemCollectionMetrics is different.")
	//}
	if len(localGetItemOutput.Attributes) != len(testGetItemOutput.Attributes) {
		panic("The size of Attributes is different.")
	}
	sortItem := make([]string, 0)
	for k := range localGetItemOutput.Attributes {
		sortItem = append(sortItem, k)
	}
	sort.Strings(sortItem)
	for _, k := range sortItem {
		if utils.Change(localGetItemOutput.Attributes[k]) != utils.Change(testGetItemOutput.Attributes[k]) {
			panic("The Attributes is different.")
		}
	}
	log.Info(context.TODO(), "putItem is successful!")
}

func UpdateItem(localGetItemOutput, testGetItemOutput dynamodb.UpdateItemOutput) {
	//TODO: judge
	//if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
	//	panic("judgeUpdateItem test is fail: ConsumedCapacity is different.")
	//}
	// TODO:
	//if localGetItemOutput.ItemCollectionMetrics != testGetItemOutput.ItemCollectionMetrics {
	//	panic("judgeUpdateItem test is fail: ItemCollectionMetrics is different.")
	//}
	if len(localGetItemOutput.Attributes) != len(testGetItemOutput.Attributes) {
		panic("The size of Attributes is different.")
	}
	sortItem := make([]string, 0)
	for k := range localGetItemOutput.Attributes {
		sortItem = append(sortItem, k)
	}
	sort.Strings(sortItem)
	for _, k := range sortItem {
		if utils.Change(localGetItemOutput.Attributes[k]) != utils.Change(testGetItemOutput.Attributes[k]) {
			panic("The Attributes is different.")
		}
	}
	log.Info(context.TODO(), "updateItem is successful!")
}

func DeleteItem(localDeleteItemOutput, testDeleteItemOutput dynamodb.DeleteItemOutput) {
	//TODO: judge
	//if localGetItemOutput.ConsumedCapacity != testGetItemOutput.ConsumedCapacity {
	//	panic("judgeUpdateItem test is fail: ConsumedCapacity is different.")
	//}
	// TODO:
	//if localGetItemOutput.ItemCollectionMetrics != testGetItemOutput.ItemCollectionMetrics {
	//	panic("judgeUpdateItem test is fail: ItemCollectionMetrics is different.")
	//}
	//if len(localDeleteItemOutput.Attributes) != len(testDeleteItemOutput.Attributes) {
	//	panic("The size of Attributes is different.")
	//}
	//if *localDeleteItemOutput.ConsumedCapacity.TableName != *testDeleteItemOutput.ConsumedCapacity.TableName {
	//	panic("The tableName is different.")
	//}
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

package judge

import (
	"context"
	"double_test_client/log"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetItem(localGetItemOutput dynamodb.GetItemOutput, testGetItem []byte) {
	log.Info(context.TODO(), "testGetItem([]byte):", testGetItem)
	// change
	testGetItemOutput := &dynamodb.GetItemOutput{}
	if err := json.Unmarshal(testGetItem, &testGetItemOutput); err != nil {
		log.Info(context.TODO(), "getItemOutputChageError: ", err)
		panic(err)
	}

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

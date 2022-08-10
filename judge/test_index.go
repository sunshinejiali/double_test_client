package judge

import (
	"context"
	"double_test_client/log"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Query(localQueryOutput, testQueryOutput dynamodb.QueryOutput) {
	// judge
	if localQueryOutput.ConsumedCapacity != testQueryOutput.ConsumedCapacity {
		panic("Query test is fail: ConsumedCapacity is different.")
	}
	if localQueryOutput.Count != testQueryOutput.Count {
		panic("Query test is fail: Count is different.")
	}
	if localQueryOutput.ScannedCount != testQueryOutput.ScannedCount {
		panic("Query test is fail: ScannedCountv is different.")
	}
	if len(localQueryOutput.Items) != len(testQueryOutput.Items) {
		panic("Query test is fail: the size of items is different.")
	}
	for index, localMap := range localQueryOutput.Items {
		testMap := testQueryOutput.Items[index]
		if len(localMap) != len(testMap) {
			panic("The size of item is different.")
		}
		for key, value := range localMap {
			if value != testMap[key] {
				panic("The value is different.")
			}
		}

	}
	// result
	log.Info(context.TODO(), "query is successful!")
}

func Scan(localScanOutput, testScanOutput dynamodb.ScanOutput) {
	// judge
	if localScanOutput.ConsumedCapacity != testScanOutput.ConsumedCapacity {
		panic("Scan test is fail: ConsumedCapacity is different.")
	}
	if localScanOutput.Count != testScanOutput.Count {
		panic("Scan test is fail: Count is different.")
	}
	if localScanOutput.ScannedCount != testScanOutput.ScannedCount {
		panic("Scan test is fail: ScannedCountv is different.")
	}
	if len(localScanOutput.Items) != len(testScanOutput.Items) {
		panic("Scan test is fail: the size of items is different.")
	}
	for index, localMap := range localScanOutput.Items {
		testMap := testScanOutput.Items[index]
		if len(localMap) != len(testMap) {
			panic("The size of item is different.")
		}
		for key, value := range localMap {
			if value != testMap[key] {
				panic("The value is different.")
			}
		}

	}
	// result
	log.Info(context.TODO(), "query is successful!")
}

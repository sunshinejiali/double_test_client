package judge

import (
	"context"
	"double_test_client/log"
	"double_test_client/utils"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sort"
)

func Query(localQueryOutput, testQueryOutput dynamodb.QueryOutput) {
	// TODO: judge ConsumedCapacity
	//if localQueryOutput.ConsumedCapacity != testQueryOutput.ConsumedCapacity {
	//	panic("Query test is fail: ConsumedCapacity is different.")
	//}
	if *localQueryOutput.Count != *testQueryOutput.Count {
		panic("Query test is fail: Count is different.")
	}
	if *localQueryOutput.ScannedCount != *testQueryOutput.ScannedCount {
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
			if utils.Change(value) != utils.Change(testMap[key]) {
				panic("The value is different.")
			}
		}

	}
	// result
	log.Info(context.TODO(), "query is successful!")
}

func Scan(localScanOutput, testScanOutput dynamodb.ScanOutput) {
	// TODO: judge ConsumedCapacity
	//if localScanOutput.ConsumedCapacity != testScanOutput.ConsumedCapacity {
	//	panic("Scan test is fail: ConsumedCapacity is different.")
	//}
	if *localScanOutput.Count != *testScanOutput.Count {
		panic("Scan test is fail: Count is different.")
	}
	if *localScanOutput.ScannedCount != *testScanOutput.ScannedCount {
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
		sortItem := make([]string, 0)
		for k := range testMap {
			sortItem = append(sortItem, k)
		}
		sort.Strings(sortItem)
		for _, k := range sortItem {
			if utils.Change(testMap[k]) != utils.Change(localMap[k]) {
				panic("The Attributes is different.")
			}
		}
	}
	// result
	log.Info(context.TODO(), "scan is successful!")
}

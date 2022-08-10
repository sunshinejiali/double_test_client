package judge

import (
	"context"
	"double_test_client/log"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
CreateTable

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

attributeDefinitions
tableName
keySchema
tableStatus
//CreationDateTime
ProvisionedThroughput
//TbaleSizeBytes
//ItemCount
//tableArn
*/
func CreateTable(localCreateTableOutput dynamodb.CreateTableOutput, testCreateTable []byte) {
	log.Info(context.TODO(), "testCreateTable([]byte):", testCreateTable)
	// change
	testCreateTablesOutput := &dynamodb.CreateTableOutput{}
	if err := json.Unmarshal(testCreateTable, &testCreateTablesOutput); err != nil {
		log.Info(context.TODO(), "createTableOutputChageError: ", err)
		panic(err)
	}
	localDescributeTable := localCreateTableOutput.TableDescription
	testDescributetable := testCreateTablesOutput.TableDescription
	// judge
	if localDescributeTable.TableName != testDescributetable.TableName {
		panic("judgeCreateTable test is fail: TableName is different.")
	}
	if localDescributeTable.TableStatus != testDescributetable.TableStatus {
		panic("judgeCreateTable test is fail: TableStatus is different.")
	}
	if len(localDescributeTable.AttributeDefinitions) != len(testDescributetable.AttributeDefinitions) {
		panic("The size of AttributeDefinitions is different.")
	}
	for key, name := range localDescributeTable.AttributeDefinitions {
		if name != testDescributetable.AttributeDefinitions[key] {
			panic("The AttributeDefinitions is different.")
		}
	}
	if len(localDescributeTable.KeySchema) != len(testDescributetable.KeySchema) {
		panic("The size of KeySchema is different.")
	}
	for key, name := range localDescributeTable.KeySchema {
		if name != testDescributetable.KeySchema[key] {
			panic("The KeySchema is different.")
		}
	}
	if (localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	// other attributes:
	//if localDescributeTable.ItemCount != testDescributetable.ItemCount {
	//	panic("judgeCreateTable test is fail: ItemCount is different.")
	//}
	//if localDescributeTable.TableSizeBytes != testDescributetable.TableSizeBytes {
	//	panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	//}

	log.Info(context.TODO(), "createTable is successful!")
}

func DeleteTable(localDeleteTableOutput dynamodb.DeleteTableOutput, testDeleteTable []byte) {
	log.Info(context.TODO(), "testDeleteTable([]byte):", testDeleteTable)
	// change
	testDeleteTableOutput := &dynamodb.DeleteTableOutput{}
	if err := json.Unmarshal(testDeleteTable, &testDeleteTableOutput); err != nil {
		log.Info(context.TODO(), "deleteTableOutputChageError: ", err)
		panic(err)
	}
	localDescributeTable := localDeleteTableOutput.TableDescription
	testDescributetable := testDeleteTableOutput.TableDescription
	// judge
	if localDescributeTable.TableName != testDescributetable.TableName {
		panic("judgeDeleteTable test is fail: TableName is different.")
	}
	if localDescributeTable.TableStatus != testDescributetable.TableStatus {
		panic("judgeDeleteTable test is fail: TableStatus is different.")
	}
	if len(localDescributeTable.AttributeDefinitions) != len(testDescributetable.AttributeDefinitions) {
		panic("The size of AttributeDefinitions is different.")
	}
	for key, name := range localDescributeTable.AttributeDefinitions {
		if name != testDescributetable.AttributeDefinitions[key] {
			panic("The AttributeDefinitions is different.")
		}
	}
	if len(localDescributeTable.KeySchema) != len(testDescributetable.KeySchema) {
		panic("The size of KeySchema is different.")
	}
	for key, name := range localDescributeTable.KeySchema {
		if name != testDescributetable.KeySchema[key] {
			panic("The KeySchema is different.")
		}
	}
	if (localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	// other attributes:
	//if localDescributeTable.ItemCount != testDescributetable.ItemCount {
	//	panic("judgeCreateTable test is fail: ItemCount is different.")
	//}
	//if localDescributeTable.TableSizeBytes != testDescributetable.TableSizeBytes {
	//	panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	//}

	log.Info(context.TODO(), "deleteTable is successful!")
}

func UpdateTable(localUpdateTableOutput dynamodb.UpdateTableOutput, testUpdateTable []byte) {
	log.Info(context.TODO(), "testUpdateTable([]byte):", testUpdateTable)
	// change
	testUpdateTableOutput := &dynamodb.UpdateTableOutput{}
	if err := json.Unmarshal(testUpdateTable, &testUpdateTableOutput); err != nil {
		log.Info(context.TODO(), "updateTableOutputChageError: ", err)
		panic(err)
	}
	localDescributeTable := localUpdateTableOutput.TableDescription
	testDescributetable := testUpdateTableOutput.TableDescription
	// judge
	if localDescributeTable.TableName != testDescributetable.TableName {
		panic("judgeUpdateTable test is fail: TableName is different.")
	}
	if localDescributeTable.TableStatus != testDescributetable.TableStatus {
		panic("judgeUpdateTable test is fail: TableStatus is different.")
	}
	if len(localDescributeTable.AttributeDefinitions) != len(testDescributetable.AttributeDefinitions) {
		panic("The size of AttributeDefinitions is different.")
	}
	for key, name := range localDescributeTable.AttributeDefinitions {
		if name != testDescributetable.AttributeDefinitions[key] {
			panic("The AttributeDefinitions is different.")
		}
	}
	if len(localDescributeTable.KeySchema) != len(testDescributetable.KeySchema) {
		panic("The size of KeySchema is different.")
	}
	for key, name := range localDescributeTable.KeySchema {
		if name != testDescributetable.KeySchema[key] {
			panic("The KeySchema is different.")
		}
	}
	if (localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	// other attributes:
	//if localDescributeTable.ItemCount != testDescributetable.ItemCount {
	//	panic("judgeCreateTable test is fail: ItemCount is different.")
	//}
	//if localDescributeTable.TableSizeBytes != testDescributetable.TableSizeBytes {
	//	panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	//}

	log.Info(context.TODO(), "updateTable is successful!")
}

func DescribeTable(localDescribeTableOutput dynamodb.DescribeTableOutput, testDescribeTable []byte) {
	log.Info(context.TODO(), "testDescribeTable([]byte):", testDescribeTable)
	// change
	testDescribeTableOutput := &dynamodb.DescribeTableOutput{}
	if err := json.Unmarshal(testDescribeTable, &testDescribeTableOutput); err != nil {
		log.Info(context.TODO(), "describeTableOutputChageError: ", err)
		panic(err)
	}
	localDescributeTable := localDescribeTableOutput.Table
	testDescributetable := testDescribeTableOutput.Table
	// judge
	if localDescributeTable.TableName != testDescributetable.TableName {
		panic("judgeDescribeTable test is fail: TableName is different.")
	}
	if localDescributeTable.TableStatus != testDescributetable.TableStatus {
		panic("judgeDescribeTable test is fail: TableStatus is different.")
	}
	if len(localDescributeTable.AttributeDefinitions) != len(testDescributetable.AttributeDefinitions) {
		panic("The size of AttributeDefinitions is different.")
	}
	for key, name := range localDescributeTable.AttributeDefinitions {
		if name != testDescributetable.AttributeDefinitions[key] {
			panic("The AttributeDefinitions is different.")
		}
	}
	if len(localDescributeTable.KeySchema) != len(testDescributetable.KeySchema) {
		panic("The size of KeySchema is different.")
	}
	for key, name := range localDescributeTable.KeySchema {
		if name != testDescributetable.KeySchema[key] {
			panic("The KeySchema is different.")
		}
	}
	if (localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	// other attributes:
	//if localDescributeTable.ItemCount != testDescributetable.ItemCount {
	//	panic("judgeCreateTable test is fail: ItemCount is different.")
	//}
	//if localDescributeTable.TableSizeBytes != testDescributetable.TableSizeBytes {
	//	panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	//}

	log.Info(context.TODO(), "describeTable is successful!")
}

/*
ListTables

	tableNames
	lastEvaluatedTableName
*/
func ListTables(localListTablesOutput dynamodb.ListTablesOutput, testListTables []byte) {
	log.Info(context.TODO(), "testListTables([]byte):", testListTables)
	// change
	testListTablesOutput := &dynamodb.ListTablesOutput{}
	if err := json.Unmarshal(testListTables, &testListTablesOutput); err != nil {
		log.Info(context.TODO(), "listTablesOutputChageError: ", err)
		panic(err)
	}
	// judge
	if localListTablesOutput.LastEvaluatedTableName != testListTablesOutput.LastEvaluatedTableName {
		panic("judgeListTables test is fail: LastEvaluatedTableName is different.")
	}
	if len(localListTablesOutput.TableNames) != len(testListTablesOutput.TableNames) {
		panic("The size of tableNames is different.")
	}
	for index, name := range localListTablesOutput.TableNames {
		if name != testListTablesOutput.TableNames[index] {
			panic("The tableName is different.")
		}
	}
	// result
	log.Info(context.TODO(), "listTables is successful!")
}

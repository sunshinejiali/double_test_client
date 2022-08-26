package judge

import (
	"context"
	"double_test_client/log"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sort"
)

type IndexG []dynamodb.GlobalSecondaryIndexDescription

func (s IndexG) Len() int {
	return len(s)
}

func (s IndexG) Less(i, j int) bool {
	return *s[i].IndexName < *s[j].IndexName
}

func (s IndexG) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type IndexL []dynamodb.LocalSecondaryIndexDescription

func (s IndexL) Len() int {
	return len(s)
}

func (s IndexL) Less(i, j int) bool {
	return *s[i].IndexName < *s[j].IndexName
}

func (s IndexL) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func CreateTable(localCreateTableOutput, testCreateTablesOutput dynamodb.CreateTableOutput) {
	localDescributeTable := localCreateTableOutput.TableDescription
	testDescributetable := testCreateTablesOutput.TableDescription
	if *localDescributeTable.TableName != *testDescributetable.TableName {
		panic("judgeCreateTable test is fail: TableName is different.")
	}
	if *localDescributeTable.TableArn != *testDescributetable.TableArn {
		panic("judgeCreateTable test is fail: TableArn is different.")
	}
	if *localDescributeTable.TableSizeBytes != *testDescributetable.TableSizeBytes {
		panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	}
	if *localDescributeTable.ItemCount != *testDescributetable.ItemCount {
		panic("judgeCreateTable test is fail: ItemCount is different.")
	}

	if *localDescributeTable.TableStatus != *testDescributetable.TableStatus {
		panic("judgeCreateTable test is fail: TableStatus is different.")
	}
	judgeAttribute(localDescributeTable.AttributeDefinitions, testDescributetable.AttributeDefinitions)

	judgeSchema(localDescributeTable.KeySchema, testDescributetable.KeySchema)

	if (*localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != *testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (*localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != *localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	if localDescributeTable.GlobalSecondaryIndexes != nil {
		if len(localDescributeTable.GlobalSecondaryIndexes) != len(testDescributetable.GlobalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeGlobal(localDescributeTable.GlobalSecondaryIndexes, testDescributetable.GlobalSecondaryIndexes)
	}

	if localDescributeTable.LocalSecondaryIndexes != nil {
		if len(localDescributeTable.LocalSecondaryIndexes) != len(testDescributetable.LocalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeLocal(localDescributeTable.LocalSecondaryIndexes, testDescributetable.LocalSecondaryIndexes)
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

func DeleteTable(localDeleteTableOutput, testDeleteTableOutput dynamodb.DeleteTableOutput) {
	localDescributeTable := localDeleteTableOutput.TableDescription
	testDescributetable := testDeleteTableOutput.TableDescription

	if *localDescributeTable.TableName != *testDescributetable.TableName {
		panic("judgeDeleteTable test is fail: TableName is different.")
	}
	if *localDescributeTable.TableStatus != *testDescributetable.TableStatus {
		panic("judgeDeleteTable test is fail: TableStatus is different.")
	}
	if *localDescributeTable.TableArn != *testDescributetable.TableArn {
		panic("judgeDeleteTable test is fail: TableArn is different.")
	}
	// TODO: tablesizebytes is  different
	//if *localDescributeTable.TableSizeBytes != *testDescributetable.TableSizeBytes {
	//	panic("judgeDeleteTable test is fail: TableSizeBytes is different.")
	//}
	if *localDescributeTable.ItemCount != *testDescributetable.ItemCount {
		panic("judgeDeleteTable test is fail: ItemCount is different.")
	}
	judgeAttribute(localDescributeTable.AttributeDefinitions, testDescributetable.AttributeDefinitions)
	judgeSchema(localDescributeTable.KeySchema, testDescributetable.KeySchema)

	if (*localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != *testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (*localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != *localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	if localDescributeTable.GlobalSecondaryIndexes != nil {
		if len(localDescributeTable.GlobalSecondaryIndexes) != len(testDescributetable.GlobalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeGlobal(localDescributeTable.GlobalSecondaryIndexes, testDescributetable.GlobalSecondaryIndexes)
	}

	if localDescributeTable.LocalSecondaryIndexes != nil {
		if len(localDescributeTable.LocalSecondaryIndexes) != len(testDescributetable.LocalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeLocal(localDescributeTable.LocalSecondaryIndexes, testDescributetable.LocalSecondaryIndexes)
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

func UpdateTable(localUpdateTableOutput, testUpdateTableOutput dynamodb.UpdateTableOutput) {
	localDescributeTable := localUpdateTableOutput.TableDescription
	testDescributetable := testUpdateTableOutput.TableDescription

	if *localDescributeTable.TableName != *testDescributetable.TableName {
		panic("judgeUpdateTable test is fail: TableName is different.")
	}
	if *localDescributeTable.TableStatus != *testDescributetable.TableStatus {
		panic("judgeUpdateTable test is fail: TableStatus is different.")
	}
	if *localDescributeTable.TableArn != *testDescributetable.TableArn {
		panic("judgeUpdateTable test is fail: TableArn is different.")
	}
	//TODO:
	//if *localDescributeTable.TableSizeBytes != *testDescributetable.TableSizeBytes {
	//	panic("judgeUpdateTable test is fail: TableSizeBytes is different.")
	//}
	if *localDescributeTable.ItemCount != *testDescributetable.ItemCount {
		panic("judgeUpdateTable test is fail: ItemCount is different.")
	}
	judgeAttribute(localDescributeTable.AttributeDefinitions, testDescributetable.AttributeDefinitions)
	judgeSchema(localDescributeTable.KeySchema, testDescributetable.KeySchema)

	// other attributes:
	//if localDescributeTable.ItemCount != testDescributetable.ItemCount {
	//	panic("judgeCreateTable test is fail: ItemCount is different.")
	//}
	//if localDescributeTable.TableSizeBytes != testDescributetable.TableSizeBytes {
	//	panic("judgeCreateTable test is fail: TableSizeBytes is different.")
	//}
	if (*localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != *testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (*localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != *localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}

	log.Info(context.TODO(), "updateTable is successful!")
}

func DescribeTable(localDescribeTableOutput, testDescribeTableOutput dynamodb.DescribeTableOutput) {
	localDescributeTable := localDescribeTableOutput.Table
	testDescributetable := testDescribeTableOutput.Table
	judgeName(localDescributeTable.TableName, testDescributetable.TableName)
	if *localDescributeTable.TableStatus != *testDescributetable.TableStatus {
		panic("judgeDescribeTable test is fail: TableStatus is different.")
	}
	judgeAttribute(localDescributeTable.AttributeDefinitions, testDescributetable.AttributeDefinitions)
	judgeSchema(localDescributeTable.KeySchema, testDescributetable.KeySchema)
	if (*localDescributeTable.ProvisionedThroughput.ReadCapacityUnits != *testDescributetable.ProvisionedThroughput.ReadCapacityUnits) || (*localDescributeTable.ProvisionedThroughput.WriteCapacityUnits != *localDescributeTable.ProvisionedThroughput.WriteCapacityUnits) {
		panic("The size of ProvisionedThroughput is different.")
	}
	if *localDescributeTable.TableArn != *testDescributetable.TableArn {
		panic("judgeDescribeTable test is fail: TableArn is different.")
	}
	//TODO:
	//if *localDescributeTable.TableSizeBytes != *testDescributetable.TableSizeBytes {
	//	panic("judgeDescribeTable test is fail: TableSizeBytes is different.")
	//}
	if *localDescributeTable.ItemCount != *testDescributetable.ItemCount {
		panic("judgeDescribeTable test is fail: ItemCount is different.")
	}

	if localDescributeTable.GlobalSecondaryIndexes != nil {
		if len(localDescributeTable.GlobalSecondaryIndexes) != len(testDescributetable.GlobalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeGlobal(localDescributeTable.GlobalSecondaryIndexes, testDescributetable.GlobalSecondaryIndexes)
	}

	if localDescributeTable.LocalSecondaryIndexes != nil {
		if len(localDescributeTable.LocalSecondaryIndexes) != len(testDescributetable.LocalSecondaryIndexes) {
			panic("The size of globalIndex is different.")
		}
		//judgeLocal(localDescributeTable.LocalSecondaryIndexes, testDescributetable.LocalSecondaryIndexes)
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

func ListTables(localListTablesOutput, testListTablesOutput dynamodb.ListTablesOutput) {

	judgeName(localListTablesOutput.LastEvaluatedTableName, testListTablesOutput.LastEvaluatedTableName)

	if len(localListTablesOutput.TableNames) != len(testListTablesOutput.TableNames) {
		panic("The size of tableNames is different.")
	}
	if len(localListTablesOutput.TableNames) != 0 {
		for index, name := range localListTablesOutput.TableNames {
			if *name != *testListTablesOutput.TableNames[index] {
				panic("The tableName is different.")
			}
		}
	}
	log.Info(context.TODO(), "listTables is successful!")
}

func judgeAttribute(local, test []*dynamodb.AttributeDefinition) {
	if len(local) != len(test) {
		panic("The size of AttributeDefinitions is different.")
	}
	for key, name := range local {
		if (*name.AttributeName != *test[key].AttributeName) || (*name.AttributeType != *test[key].AttributeType) {
			panic("The AttributeDefinitions is different.")
		}
	}
}

func judgeSchema(local, test []*dynamodb.KeySchemaElement) {
	if len(local) != len(test) {
		panic("The size of KeySchema is different.")
	}
	for key, name := range local {
		if (*name.AttributeName != *test[key].AttributeName) || (*name.KeyType != *test[key].KeyType) {
			panic("The KeySchema is different.")
		}
	}
}

func judgeName(localName, testName *string) {
	if (localName == nil && testName != nil) || (localName != nil && testName == nil) {
		panic("Test is fail: TableName is different.")
	} else if localName == nil && testName == nil {

	} else if *localName != *testName {
		fmt.Println(*localName, *testName, "**************")
		panic("Test is fail: TableName is different.")
	}
}

func judgeGlobal(localDescribeTable, testDescribeTable []*dynamodb.GlobalSecondaryIndexDescription) {
	des := make(IndexG, 0)
	for _, v := range localDescribeTable {
		x := dynamodb.GlobalSecondaryIndexDescription{
			IndexArn:       v.IndexArn,
			IndexName:      v.IndexName,
			IndexSizeBytes: v.IndexSizeBytes,
			ItemCount:      v.ItemCount,
			KeySchema:      v.KeySchema,
			Projection:     v.Projection,
		}
		des = append(des, x)
	}
	sort.Sort(des)
	desL := make(IndexG, 0)
	for _, v := range testDescribeTable {
		x := dynamodb.GlobalSecondaryIndexDescription{
			IndexArn:       v.IndexArn,
			IndexName:      v.IndexName,
			IndexSizeBytes: v.IndexSizeBytes,
			ItemCount:      v.ItemCount,
			KeySchema:      v.KeySchema,
			Projection:     v.Projection,
		}
		desL = append(desL, x)
	}
	sort.Sort(desL)
	for index, local := range localDescribeTable {
		fmt.Println(index, *local.IndexStatus, *desL[index].IndexStatus)
		judgeName(local.IndexName, desL[index].IndexName)
		judgeSchema(local.KeySchema, desL[index].KeySchema)
		if *local.IndexStatus != *desL[index].IndexStatus {
			panic("judgeDescribeTable test is fail: IndexStatus is different.")
		}
		if local.ItemCount != desL[index].ItemCount {
			panic("judgeDescribeTable test is fail: ItemCount is different.")
		}
		if *local.IndexArn != *desL[index].IndexArn {
			panic("judgeCreateTable test is fail: IndexArn is different.")
		}
		if (*local.ProvisionedThroughput.ReadCapacityUnits != *desL[index].ProvisionedThroughput.ReadCapacityUnits) || (*local.ProvisionedThroughput.WriteCapacityUnits != *desL[index].ProvisionedThroughput.WriteCapacityUnits) {
			panic("The size of ProvisionedThroughput is different.")
		}
		judgeProjection(local.Projection, desL[index].Projection)
	}
}

func judgeLocal(localDescribeTable, testDescribeTable []*dynamodb.LocalSecondaryIndexDescription) {
	des := make(IndexL, 0)
	for _, v := range localDescribeTable {
		x := dynamodb.LocalSecondaryIndexDescription{
			IndexArn:       v.IndexArn,
			IndexName:      v.IndexName,
			IndexSizeBytes: v.IndexSizeBytes,
			ItemCount:      v.ItemCount,
			KeySchema:      v.KeySchema,
			Projection:     v.Projection,
		}
		des = append(des, x)
	}
	sort.Sort(des)
	desL := make(IndexL, 0)
	for _, v := range testDescribeTable {
		x := dynamodb.LocalSecondaryIndexDescription{
			IndexArn:       v.IndexArn,
			IndexName:      v.IndexName,
			IndexSizeBytes: v.IndexSizeBytes,
			ItemCount:      v.ItemCount,
			KeySchema:      v.KeySchema,
			Projection:     v.Projection,
		}
		desL = append(desL, x)
	}
	sort.Sort(desL)
	for index, local := range localDescribeTable {
		judgeName(local.IndexName, desL[index].IndexName)
		judgeSchema(local.KeySchema, desL[index].KeySchema)
		if *local.ItemCount != *desL[index].ItemCount {
			panic("judgeDescribeTable test is fail: ItemCount is different.")
		}
		if *local.IndexArn != *desL[index].IndexArn {
			panic("judgeCreateTable test is fail: IndexArn is different.")
		}
		judgeProjection(local.Projection, desL[index].Projection)
	}
}

func judgeProjection(local, test *dynamodb.Projection) {
	if *local.ProjectionType != *test.ProjectionType {
		panic("The size of Projection is different.")
	}
	if len(local.NonKeyAttributes) != len(test.NonKeyAttributes) {
		panic("The size of Projection is different.")
	}
	for index, name := range local.NonKeyAttributes {
		if *name != *test.NonKeyAttributes[index] {
			panic("The name of Projection is different.")
		}
	}
}

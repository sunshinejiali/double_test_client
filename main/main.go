package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//deleteTable(router)
	//
	//listTables(router)
	//
	//createTableTest1(router)
	//
	//listTables(router)
	//
	//describeTable(router)
	//
	////updateTable(router)
	//updateTable2(router)
	//
	//describeTable(router)
	//
	//listTables(router)
	//
	//log.Infof(context.TODO(), "waiting...")
	//time.Sleep(10000000000)
	//log.Infof(context.TODO(), "===============")
	//
	//describeTable(router)
	//
	//deleteTable(router)
	////  ====== item =====

	deleteTable(router)

	createTableTest2(router)

	putItem2(router)
	//

	describeTable(router)

	getItem2(router)
	//
	deleteItem(router)

	describeTable(router)
	//
	putItem2(router)

	describeTable(router)
	//
	updateItem2(router)
	//
	////  ====== index =====

	scan(router)

	query(router)

	deleteTable(router)

}

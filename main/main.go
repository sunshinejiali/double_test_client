package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//deleteTable(router)

	listTables(router)

	listTables(router)

	createTableTest3(router)

	listTables(router)

	describeTable(router)

	//updateTable(router)
	updateTable4(router)

	describeTable(router)

	listTables(router)
	//
	deleteTable(router)
	//////  ====== item =====
	//putItem(router)
	////
	//getItem(router)
	////
	//deleteItem(router)
	////
	//putItem(router)
	////
	//updateItem(router)
	////
	//////  ====== index =====
	////query(router)
	//scan(router)
	//deleteTable(router)

	//batchWriteItem(router)
	//
	//batchGetItem(router)

}

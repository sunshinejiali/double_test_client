package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	listTables(router)

	createtable(router)

	listTables(router)

	describeTable(router)

	listTables(router)

	updateTable(router)

	//  ====== item =====
	putItem(router)

	getItem(router)

	deleteItem(router)

	putItem(router)

	updateItem(router)

	//  ====== index =====
	//query(router)
	scan(router)
	deleteTable(router)

	//batchWriteItem(router)

	//batchGetItem(router)

}

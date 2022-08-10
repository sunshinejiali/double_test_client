package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	//  ====== table =====
	createtable(router)

	updateTable(router)

	describeTable(router)

	listTables(router)

	deleteTable(router)

	//  ====== index =====
	//query()

	//scan()

	//  ====== item =====
	//batchGetItem()

	//batchWriteItem()

	getItem(router)

	//putItem()

	//updateItem()

	router.Run("https://localhost:8989")
}

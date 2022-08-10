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

	//  ====== item =====
	getItem(router)

	putItem(router)

	updateItem(router)

	//batchWriteItem(router)

	//batchGetItem(router)

	//  ====== index =====
	query(router)

	scan(router)

	router.Run("https://localhost:8989")
}

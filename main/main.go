package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//  ====== table =====
	createtable(router)

	//fmt.Println("======1======")

	//listTables(router)
	//
	//fmt.Println("======2======")
	//
	//describeTable(router)
	//
	//fmt.Println("======3======")
	//
	//updateTable(router)
	//
	//fmt.Println("======4======")
	//
	//deleteTable(router)
	//
	//fmt.Println("======5======")
	//
	////  ====== item =====
	//getItem(router)
	//
	//putItem(router)
	//
	//updateItem(router)
	//
	////batchWriteItem(router)
	//
	////batchGetItem(router)
	//
	////  ====== index =====
	//query(router)
	//
	//scan(router)

	router.Run("https://localhost:8989")
}

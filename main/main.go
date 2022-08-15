package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//  ====== table =====

	listTables(router)

	createtable(router)

	fmt.Println("======1======")

	listTables(router)

	fmt.Println("======2======")

	describeTable(router)

	fmt.Println("************")

	listTables(router)

	fmt.Println("======3======")

	updateTable(router)

	fmt.Println("======4======")

	//  ====== item =====
	putItem(router)

	getItem(router)

	updateItem(router)

	//  ====== index =====
	query(router)

	fmt.Println("======**======")
	scan(router)

	deleteTable(router)

	fmt.Println("======5======")

	//batchWriteItem(router)

	//batchGetItem(router)

}

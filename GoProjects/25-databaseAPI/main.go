package main

import (
	"fmt"

	"github.com/akscodebay/databaseapi/dao"
	"github.com/akscodebay/databaseapi/router"
)

func main() {
	fmt.Println("Starting server on :8080")
	dao.CreateTable()
	router.GenerateRoutes()
	defer fmt.Println("Server stopped")
}

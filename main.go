package main

import (
	"munenendereba/go-crud-api/db"
	"munenendereba/go-crud-api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}

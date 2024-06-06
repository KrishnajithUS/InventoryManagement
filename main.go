package main

import (
	"InventoryManagement/database"
	"InventoryManagement/server"
)

func main() {
	database.Init()
	server.Init()
}

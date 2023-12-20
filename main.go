package main

import (
	"github.com/JeanIzar/crud-go/src/config"
	"github.com/JeanIzar/crud-go/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//run all routes
	routes.Routes()
}

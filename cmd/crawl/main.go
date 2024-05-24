package main

import (
	"github.com/shibu1x/ur_v3/pkg/db"
	"github.com/shibu1x/ur_v3/pkg/models"
)

func main() {
	// データベース接続を確立
	database := db.ConnectDB()
	defer database.Close()

	models.UpdateHousesAll()
}

package main

import (
	"backend/config"
	"backend/models"
)

func main() {
	db := config.InitDB()

	models.InitDB(db)
}

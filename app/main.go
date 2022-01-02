package main

import (
	"app/models"
	"app/routes"
)


func main() {
	
	db := models.SetupDB()
	
	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.Activity{})

	r := routes.SetupRoutes(db)
	r.Run(":3030")
}
package main

import (
	"net/http"

	"phonebook/controllers"
	"phonebook/database"
	"phonebook/logger"
	"phonebook/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Init()
	logger.Log.Info("Application started")

	r := gin.Default()

	db, err := database.Connect()
	if err != nil {
		logger.Log.Error("Failed to connect to the database: %v", err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	contactController := controllers.NewContactController(db)
	routes.SetupRoutes(r, contactController)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	gin.SetMode(gin.DebugMode)

	r.Run(":8080")
}

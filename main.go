package main

import (
	"PersonalScheduleAPI/db"
	"PersonalScheduleAPI/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.InitDatabase()
	defer database.DB.Close()

	router := gin.Default()
	router.GET("/schedule", handlers.GetScheduleItems)
	router.POST("/schedule", handlers.CreateScheduleItem)
	router.GET("/schedule/:id", handlers.GetScheduleItem)
	router.PUT("/schedule/:id", handlers.UpdateScheduleItem)
	router.DELETE("/schedule/:id", handlers.DeleteScheduleItem)
	log.Println("Server started at :8080")
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

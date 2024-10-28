package handlers

import (
	database "PersonalScheduleAPI/db"
	"PersonalScheduleAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get all schedule items
func GetScheduleItems(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, start_time, end_time, status FROM schedule")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.ScheduleItem
	for rows.Next() {
		var item models.ScheduleItem
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.StartTime, &item.EndTime, &item.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

// Add a new schedule item
func CreateScheduleItem(c *gin.Context) {
	var item models.ScheduleItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	query := "INSERT INTO schedule (title, description, start_time, end_time, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := database.DB.QueryRow(query, item.Title, item.Description, item.StartTime, item.EndTime, item.Status).Scan(&item.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// Get a single schedule item by ID
func GetScheduleItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var item models.ScheduleItem
	query := "SELECT id, title, description, start_time, end_time, status FROM schedule WHERE id = $1"
	err = database.DB.QueryRow(query, id).Scan(&item.ID, &item.Title, &item.Description, &item.StartTime, &item.EndTime, &item.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Update an existing schedule item
func UpdateScheduleItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var item models.ScheduleItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	query := "UPDATE schedule SET title = $1, description = $2, start_time = $3, end_time = $4, status = $5 WHERE id = $6"
	_, err = database.DB.Exec(query, item.Title, item.Description, item.StartTime, item.EndTime, item.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Delete a schedule item
func DeleteScheduleItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	query := "DELETE FROM schedule WHERE id = $1"
	_, err = database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

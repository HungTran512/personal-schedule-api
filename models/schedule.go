package models

import "time"

// ScheduleItem represents a personal schedule item
type ScheduleItem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"` // e.g., "planned", "completed"
}

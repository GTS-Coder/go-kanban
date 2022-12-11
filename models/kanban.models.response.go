package models

import (
	"time"
)

// 👈 DBResponse is used to get data from database
type KanbanResponse struct {
	Board struct {
		Cards []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Assignee    []struct {
				ID     string `json:"id"`
				Avatar string `json:"avatar"`
				Name   string `json:"name"`
			} `json:"assignee"`
			Due         []int64       `json:"due"`
			Attachments []interface{} `json:"attachments"`
			Comments    []struct {
				ID          string    `json:"id"`
				Avatar      string    `json:"avatar"`
				Name        string    `json:"name"`
				CreatedAt   time.Time `json:"createdAt"`
				MessageType string    `json:"messageType"`
				Message     string    `json:"message"`
			} `json:"comments"`
			Completed bool `json:"completed"`
		} `json:"cards"`
		Columns []struct {
			ID      string   `json:"id"`
			Name    string   `json:"name"`
			CardIds []string `json:"cardIds"`
		} `json:"columns"`
		ColumnOrder []string `json:"columnOrder"`
	} `json:"board"`
}

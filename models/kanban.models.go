package models

import (
	"time"
)

// 👈 DBResponse is used to sent data to database

type KanbanInput struct {
	Board struct {
		Cards []struct {
			ID          string `json:"id" bson:"id"`
			Name        string `json:"name" bson:"name"`
			Description string `json:"description" bson:"description"`
			Assignee    []struct {
				ID     string `json:"id" bson:"id"`
				Avatar string `json:"avatar" bson:"avatar"`
				Name   string `json:"name" bson:"name"`
			} `json:"assignee" bson:"assignee"`
			Due         []int64       `json:"due" bson:"due"`
			Attachments []interface{} `json:"attachments" bson:"attachments"`
			Comments    []struct {
				ID          string    `json:"id" bson:"id"`
				Avatar      string    `json:"avatar" bson:"avatar"`
				Name        string    `json:"name" bson:"name"`
				CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
				MessageType string    `json:"messageType" bson:"messageType"`
				Message     string    `json:"message" bson:"message"`
			} `json:"comments" bson:"comments"`
			Completed bool `json:"completed" bson:"completed"`
		} `json:"cards" bson:"cards"`
		Columns []struct {
			ID      string   `json:"id" bson:"id"`
			Name    string   `json:"name" bson:"name"`
			CardIds []string `json:"cardIds" bson:"cardIds"`
		} `json:"columns" bson:"columns"`
		ColumnOrder []string `json:"columnOrder" bson:"columnOrder"`
	} `json:"board" bson:"board"`
}

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

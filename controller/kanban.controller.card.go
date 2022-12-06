package controller

import (
	"context"
	"my-kanban/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientAddTask := models.AddCardInput{}

		if err := c.ShouldBindJSON(&ClientAddTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		update := bson.M{
			"$push": bson.M{
				"board.columns.$.cardIds": ClientAddTask.Card.ID,
				"board.cards":             ClientAddTask.Card,
			},
		}

		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.columns.id": ClientAddTask.ColumnID}, update).Decode(&ClientAddTask)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "Task added successfully"})
	}
}

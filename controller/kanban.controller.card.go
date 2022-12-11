package controller

import (
	"context"
	"fmt"
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

		fmt.Println("ClientAddTask: ", ClientAddTask)

		update := bson.M{
			"$push": bson.M{
				"board.columns.$.cardIds": ClientAddTask.Card.ID,
				"board.cards":             ClientAddTask.Card,
			},
		}

		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.columns.id": ClientAddTask.ColumnID}, update).Decode(&ClientAddTask)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  ClientAddTask,
			})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "Add task success"})
	}
}

func MarkDoneTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientDoneTask := models.MarkDoneInput{}

		if err := c.ShouldBindJSON(&ClientDoneTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("ClientAddTask: ", ClientDoneTask)

		update := bson.M{
			"$set": bson.M{
				"board.cards.$.completed": ClientDoneTask.Completed,
			},
		}

		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.cards.id": ClientDoneTask.CardID}, update).Decode(&ClientDoneTask)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  ClientDoneTask,
			})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "Update task success"})
	}
}

package controller

import (
	"context"
	"fmt"
	configs "my-kanban/config"
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

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientDeleteTask := models.DeleteCardInput{}
		Card := models.AddCardInput{}
		Columns := models.CloumnResponse{}

		if err := c.ShouldBindJSON(&ClientDeleteTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		deleteCard := bson.M{
			"$pull": bson.M{
				"board.cards": ClientDeleteTask.Cards,
			},
		}

		deleteCardInColumns := bson.M{
			"$pull": bson.M{
				"board.columns.$.cardIds": bson.M{"$in": []string{ClientDeleteTask.Cards.ID}},
			},
		}

		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.cards.id": ClientDeleteTask.Cards.ID}, deleteCard).Decode(&Card)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_find_card": err.Error(),
				"data":            ClientDeleteTask,
			})
			defer cancel()
			return
		}
		err = kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.columns.id": ClientDeleteTask.ColumnID}, deleteCardInColumns).Decode(&Columns)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_find_column": err.Error(),
				"data":              ClientDeleteTask,
			})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "Delete task success"})
	}
}

func AddNewColumn() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientAddNewColumns := models.AddNewColumnInput{}
		ClientGetBoard := models.DBResponse{}

		if err := c.ShouldBindJSON(&ClientAddNewColumns); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		update := bson.M{
			"$push": bson.M{
				"board.columns":     ClientAddNewColumns.Column,
				"board.columnOrder": ClientAddNewColumns.Column.ID,
			},
		}
		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"id_kanban": configs.GetEnvName("ID_KANBAN")}, update).Decode(&ClientGetBoard)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  ClientAddNewColumns,
			})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"data": ClientAddNewColumns.Column, "message": "Add new column success"})
	}
}

func AddAttachments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var option = c.Param("option")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientAddAttachments := models.AddAttachmentInput{}
		ClientGetBoard := models.DBResponse{}

		if err := c.ShouldBindJSON(&ClientAddAttachments); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		var update bson.M

		if option == "new" {

			update = bson.M{
				"$push": bson.M{
					"board.cards.$.attachments": ClientAddAttachments.Data.URL,
				},
			}
		} else if option == "delete" {
			update = bson.M{
				"$pull": bson.M{
					"board.cards.$.attachments": ClientAddAttachments.Data.URL,
				},
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "option not found",
				"data":  ClientAddAttachments,
			})
			defer cancel()
			return
		}
		err := kanbanCollection.FindOneAndUpdate(context.Background(), bson.M{"board.cards.id": ClientAddAttachments.Data.ID}, update).Decode(&ClientGetBoard)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  ClientAddAttachments,
			})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"data": ClientAddAttachments.Data, "message": "Add new attachments success"})
	}
}

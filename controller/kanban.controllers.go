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
	"go.mongodb.org/mongo-driver/mongo"
)

var kanbanCollection *mongo.Collection = configs.GetCollection(configs.DB, "boards")

// func CreateKanbanBoard() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

// 		var id_owner = c.Param("board_id")
// 		var board models.Board
// 		board.IdOwner = id_owner
// 		var err error

// 		err = c.BindJSON(&board)
// 		if err != nil {
// 			defer cancel()
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		_, err = kanbanCollection.InsertOne(ctx, board)
// 		if err != nil {
// 			defer cancel()
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		defer cancel()
// 		defer c.JSON(http.StatusOK, gin.H{"message": "Board created successfully"})
// 	}
// }

func GetKanbanBoard() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

		var board models.KanbanResponse

		var err error
		// id_owner := c.Param("board_id")
		id_owner := "6387347ca92496eddbc3a110"

		err = kanbanCollection.FindOne(ctx, bson.M{"id_kanban": id_owner}).Decode(&board)
		if err != nil {
			defer cancel()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, board)
	}
}

func UpdateKanbanColumns() gin.HandlerFunc {
	return func(c *gin.Context) {

		ClientUpdateBoard := models.CloumnUpdate{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

		if err := c.BindJSON(&ClientUpdateBoard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("columns", ClientUpdateBoard.Columns)

		_, err := kanbanCollection.UpdateOne(ctx, bson.M{"id_kanban": "6387347ca92496eddbc3a110"}, bson.M{"$set": bson.M{"board.columns": ClientUpdateBoard.Columns}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "Columns updated successfully"})
	}
}

func UpdateKanbanColumnOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		ClientUpdateBoard := models.ColumnOrderUpdate{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

		if err := c.BindJSON(&ClientUpdateBoard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("columnOrder", ClientUpdateBoard.ColumnOrder)

		_, err := kanbanCollection.UpdateOne(ctx, bson.M{"id_kanban": "6387347ca92496eddbc3a110"}, bson.M{"$set": bson.M{"board.columnOrder": ClientUpdateBoard.ColumnOrder}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "ColumnOrder updated successfully"})
	}
}

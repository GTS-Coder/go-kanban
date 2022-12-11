package controller

import (
	"context"
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

		// id_owner := c.Param("board_id")
		//id_owner to object id
		err := kanbanCollection.FindOne(ctx, bson.M{"id_kanban": configs.GetEnvName("ID_KANBAN")}).Decode(&board)
		if err != nil {
			defer cancel()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, board)
	}
}

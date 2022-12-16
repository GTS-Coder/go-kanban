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

func UpdateKanbanColumns() gin.HandlerFunc {
	return func(c *gin.Context) {

		ClientUpdateBoard := models.CloumnUpdate{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

		if err := c.ShouldBindJSON(&ClientUpdateBoard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("columnOrder", ClientUpdateBoard)

		_, err := kanbanCollection.UpdateOne(ctx, bson.M{"id_kanban": "639150414de71616c9b38134"}, bson.M{"$set": bson.M{"board.columns": ClientUpdateBoard.Columns}})

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

		_, err := kanbanCollection.UpdateOne(ctx, bson.M{"id_kanban": "639150414de71616c9b38134"}, bson.M{"$set": bson.M{"board.columnOrder": ClientUpdateBoard.ColumnOrder}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		defer cancel()
		defer c.JSON(http.StatusOK, gin.H{"message": "ColumnOrder updated successfully"})
	}
}

func RenameColumnsKanban() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientUpdate := models.ColumnRename{}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //set timeout for request to database

		if err := c.BindJSON(&ClientUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("columnOrder", ClientUpdate.ColumnID)

		fitler := bson.M{"board.columns.id": ClientUpdate.ColumnID, "id_kanban": "639150414de71616c9b38134"}

		// err := kanbanCollection.FindOne(ctx, fitler).Decode(&ClientGet)

		_, err := kanbanCollection.UpdateOne(ctx, fitler, bson.M{"$set": bson.M{"board.columns.$.name": ClientUpdate.NewName}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		ClientGetResponse := models.CloumnResponse{}

		// err = kanbanCollection.FindOne(ctx, fitler).Decode(&ClientGetResponse)

		//aggregate to get the column with the new name
		pipeline := bson.A{
			bson.M{"$match": bson.M{"board.columns.id": ClientUpdate.ColumnID}},
			bson.M{"$project": bson.M{"board.columns": bson.M{"$filter": bson.M{"input": "$board.columns", "as": "column", "cond": bson.M{"$eq": bson.A{"$$column.id", ClientUpdate.ColumnID}}}}}},
		}

		//referent document aggregation:https://www.mongodb.com/docs/manual/reference/operator/aggregation/

		result, err := kanbanCollection.Aggregate(ctx, pipeline)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		var results []bson.M

		if err = result.All(ctx, &results); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		fmt.Println("results", ClientGetResponse)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		defer cancel()
		defer c.JSON(http.StatusOK, results)
	}
}

func DeleteColumn() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		ClientAddNewColumns := models.DeleteColumnInput{}
		ClientGetBoard := models.DBResponse{}

		if err := c.ShouldBindJSON(&ClientAddNewColumns); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_json": err.Error()})
			defer cancel()
			return
		}

		update := bson.M{
			"$pull": bson.M{
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

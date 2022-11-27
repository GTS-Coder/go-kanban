package routes

import (
	"my-kanban/controller"

	"github.com/gin-gonic/gin"
)

func KanbanRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("ping",
		func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	incomingRoutes.GET("api/kanban/boards/:board_id", controller.GetKanbanBoard())
	// incomingRoutes.POST("/boards/:board_id", controller.CreateKanbanBoard())
}

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
	incomingRoutes.GET("api/kanban/board/", controller.GetKanbanBoard())

	// incomingRoutes.POST("/boards/:board_id", controller.CreateKanbanBoard())
}

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/account/signup", controller.SignUp())
	incomingRoutes.POST("/api/account/login", controller.Login())
	incomingRoutes.GET("api/account/myboard", controller.GetKanbanWhenUserLogin())
}

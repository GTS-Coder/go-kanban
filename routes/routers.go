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
	incomingRoutes.POST("api/kanban/board/update/columns", controller.UpdateKanbanColumns())
	incomingRoutes.POST("api/kanban/board/update/column-order", controller.UpdateKanbanColumnOrder())
	incomingRoutes.POST("api/kanban/board/update/column/name", controller.RenameColumnsKanban())
	incomingRoutes.POST("api/kanban/board/update/card/new", controller.AddTask())
	incomingRoutes.POST("api/kanban/board/update/card/done", controller.MarkDoneTask())
	incomingRoutes.POST("api/kanban/board/update/card/delete", controller.DeleteTask())
	incomingRoutes.POST("api/kanban/board/columns/new", controller.AddNewColumn())
	incomingRoutes.POST("api/kanban/columns/delete", controller.DeleteColumn())
	//card
	incomingRoutes.POST("api/kanban/card/attachment/:option", controller.AddAttachments())
	incomingRoutes.POST("api/kanban/card/rename", controller.RenameTask())
	incomingRoutes.POST("api/kanban/card/description", controller.ChangeDescriptionTask())
	incomingRoutes.POST("api/kanban/card/add-comment", controller.AddNewComment())

	// incomingRoutes.POST("/boards/:board_id", controller.CreateKanbanBoard())
}

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/account/register", controller.SignUp())
	incomingRoutes.POST("api/account/login", controller.Login())
	incomingRoutes.POST("api/account/my-account", controller.GetUserExist())
	incomingRoutes.GET("api/users/contacts", controller.GetAllUserContact())
	incomingRoutes.POST("api/user/refresh", controller.Refresh())

	// incomingRoutes.GET("api/account/myboard", controller.GetKanbanWhenUserLogin())
}

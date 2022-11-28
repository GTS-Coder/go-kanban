package main

import (
	configs "my-kanban/config"
	"my-kanban/middleware"
	"my-kanban/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	Port := configs.GetEnvName("PORT")
	router.Use(middleware.CORSMiddleware()) //use middleware cors for all route

	router.Use(gin.Logger()) //loggin request

	routes.KanbanRoutes(router) //use routes
	if Port == "" {
		Port = "8080"
	}

	router.Run(":" + Port)

}

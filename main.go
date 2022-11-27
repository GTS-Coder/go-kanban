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

	if Port == "" {
		Port = "8080"
	}

	router.Use(gin.Logger())                //loggin request
	router.Use(middleware.CORSMiddleware()) //use middleware cors for all route
	routes.KanbanRoutes(router)             //use routes

	router.Run(":" + Port)

}

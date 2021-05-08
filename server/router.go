package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controller "go-run-python/controllers"

	"time"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	pythonController := new(controller.PythonController)

	r.GET("/python/run", pythonController.RunProgram)
	r.GET("/python/evaluate", pythonController.EvaluateProgram)

	// r.Static("/file", "/go/bin/public")

	return r
}

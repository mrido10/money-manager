package main

import (
	"log"
	"money-manager/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	route := gin.Default()
	// gin.SetMode(gin.ReleaseMode)  					//set for release mode
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	controller.Route(route)
}

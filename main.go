package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/docs"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/graceful"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/middleware"
	"log"
	"net/http"
	"time"
)

//cfgFile is configs file' path string  for configurations, this file contains all the configurations
//for configuration set this path to your own configs
var cfgFile = "./configs/project.qa.yaml"

// @title Golang Bootcamp Ending Project- Basket API
// @version 1.0

// @contact.name Yusuf BULUT
// @contact.url https://www.linkedin.com/in/yusufbu1ut/
// @contact.email yusufblt10@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	log.Printf("%s Server starting..\n", time.Now().Format("2006/01/02 15:04:05"))

	r := gin.Default()

	middleware.RegisterMiddlewares(r)

	api.RegisterHandlers(r, cfgFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//creating server on localhost:8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	graceful.ShutdownGin(srv, time.Second*1)
}

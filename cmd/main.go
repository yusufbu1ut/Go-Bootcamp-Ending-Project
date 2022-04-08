package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/graceful"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func main() {

	fmt.Printf("%s Server starting..\n", time.Now().Format("2006/01/02 15:04:05"))

	r := gin.Default()

	//cfgFile is configs file' path string  for configurations, this file contains all the configurations
	cfgFile := "./configs/project.qa.yaml"

	middleware.RegisterMiddlewares(r)

	api.RegisterHandlers(r, cfgFile)
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

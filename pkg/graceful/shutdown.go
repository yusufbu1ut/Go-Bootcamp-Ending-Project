package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//ShutdownGin runs when ending of the server it waits given timeout duration after this waiting server will be shutdown
//On this waiting it response the requests which as came the routines before shutdown
func ShutdownGin(server *http.Server, timeout time.Duration) {

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout seconds.
	select {
	case <-ctx.Done():
		log.Printf("timeout of %s seconds.", timeout.String())
	}
	log.Println("Server exiting")
}

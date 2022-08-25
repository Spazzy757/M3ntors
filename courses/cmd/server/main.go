package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spazzy757/m3ntors/courses/pkg/router"
)

const startupLog = `
   _____                                
  / ____|                               
 | |     ___  _   _ _ __ ___  ___  ___  
 | |    / _ \| | | | '__/ __|/ _ \/ __| 
 | |___| (_) | |_| | |  \__ \  __/\__ \ 
  \_____\___/ \__,_|_|  |___/\___||___/ 
`

func main() {
	// Termination Handeling
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	// create app object and get routes
	app := router.App{}
	app.GetRouter()

	// Setup Server
	addr := fmt.Sprintf("%v:%v", "0.0.0.0", "8000")
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Run Server in Goroutine to handle Graceful Shutdowns
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("Server Start Fail")
		}
	}()
	fmt.Println(startupLog)
	log.WithFields(log.Fields{
		"host": "0.0.0.0",
		"port": "8000",
	}).Info("Starting Server")

	//Graceful Shutdown
	<-termChan
	// Any Code to Gracefully Shutdown should be done here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Graceful Shutdown Failed")
	}
	log.Info("Shutting Down Gracefully")
}

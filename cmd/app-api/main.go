package main

import (
	"context"
	gins "gin-boilerplate/gin-sever"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Init gin engine
	gins.Init()

	// Init http server, set your custom http server config
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Minute,
		Handler:      gins.GinEngine,
	}
	// server started
	go func() {
		// Server listening at server.Addr
		err := server.ListenAndServe()
		if err != nil {
			// Server listen at %s failed
		}
	}()

	// quit server gracefully
	quitSignal := make(chan os.Signal, 20)
	signal.Notify(quitSignal, os.Interrupt)
	<-quitSignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// Fail to shut down server
	}
}

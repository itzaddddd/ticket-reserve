package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itzaddddd/ticket-reserve/server"
)

func main() {
	server := server.NewServer()

	server.SetConfig()
	server.SetDb()
	server.SetValidator()
	server.SetMiddleware()
	server.SetHandler()

	port := fmt.Sprintf(":%s", server.Cfg.App.Port)
	srv := http.Server{
		Addr:    port,
		Handler: server.App,
	}

	go func() {
		fmt.Println("server is running in go routine")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	if err := server.ShutdownDb(); err != nil {
		log.Fatal(err)
	}

	log.Println("server is exited")

}

package main

import (
	"context"
	"errors"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/config"
	delivery "github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/server"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/service"
	"github.com/onemgvv/WB_L2/develop/dev11/pkg/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "configs/development.json"

func main() {
	// Config init
	var c config.Config
	if err := c.Init(configPath); err != nil {
		log.Fatalf("[CONFIG INIT] | [ERROR]: %s", err.Error())
	}

	// Empty storage init
	store := storage.Storage{}

	// HTTP Server init
	services := service.NewService(service.Deps{Storage: &store})
	handler := delivery.NewHandler(&c, services)
	app := server.NewServer(&c, handler.Init())

	// Run app
	go func() {
		if err := app.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[SERVER START] || [FAILED]: %s", err.Error())
		}
	}()
	log.Printf("Application started: \n[PORT]: %s\n", c.HTTP.Port)

	// Graceful shutdown

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := app.Stop(ctx); err != nil {
		log.Fatalf("[SERVER STOP] || [FAILED]: %s", err.Error())
	}

	log.Println("Application stopped")
}

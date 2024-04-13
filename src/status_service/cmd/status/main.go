package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pinkphantasm/hieda/src/status_service/internal/app/server"
	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/config"
)

func main() {
	addr := config.New().Addr
	if addr == "" {
		err := fmt.Errorf(
			"env variable $%s must be specified to start [%s]",
			config.ENV_ADDR,
			config.ServiceName,
		)
		log.Fatal(err)
	}

	app := server.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(addr); err != nil {
			log.Fatal(err)
		}
	}()

	sig := <-c
	log.Printf("received signal %s, shutting down...\n", sig)

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}

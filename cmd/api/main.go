package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Safwanseban/hlsGo-s3/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-signalChan
		log.Println("server closecall came", oscall)
		cancel()
		<-time.After(2 * time.Second)
		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Fatal("error shutdown", err)
		}
	}()
	handlers.NewServer(app)

	if err := app.Listen(":3000"); err != nil {

		log.Fatal("error starting server")
	}

}

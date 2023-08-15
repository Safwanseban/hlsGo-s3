package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AppServer struct {
	App *fiber.App
}

func NewServer(app *fiber.App) {
	server := &AppServer{
		App: app,
	}
	server.App.Get("", server.StoreVideo)

}
func (s *AppServer) StoreVideo(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"ping": "pong",
	})
	return nil
}

package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AppHandler struct {
}

func NewServer(app *fiber.App) HandlerMethods {
	Handler := &AppHandler{}
	app.Get("", Handler.StoreVideo)
	return Handler

}
func (s *AppHandler) StoreVideo(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"ping": "pong",
	})
	return nil
}

type HandlerMethods interface {
	StoreVideo(*fiber.Ctx) error
}

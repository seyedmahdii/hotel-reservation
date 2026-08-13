package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	appv1 := app.Group("/api/v1")

	appv1.Get("/user", api.HandleGetUsers)
	appv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}

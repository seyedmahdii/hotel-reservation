package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	app.Listen(*listenAddr)
	appv1 := app.Group("/api/v1")

	appv1.Get("/foo", handleFoo)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello, World!"})
}

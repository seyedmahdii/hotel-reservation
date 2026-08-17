package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/api"
	"github.com/seyedmahdii/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const dburi = "mongodb://localhost:27017"

func main() {
	client, err := mongo.Connect(options.Client().
		ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	appv1 := app.Group("/api/v1")

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	appv1.Get("/user", userHandler.HandleGetUsers)
	appv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}

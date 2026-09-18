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

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(options.Client().
		ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New(config)
	appv1 := app.Group("/api/v1")

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	hotelHandler := api.NewHotelHandler(hotelStore, roomStore)

	appv1.Get("/users", userHandler.HandleGetUsers)
	appv1.Get("/users/:id", userHandler.HandleGetUser)
	appv1.Post("/users", userHandler.HandlePostUser)
	appv1.Delete("/users/:id", userHandler.HandleDeleteUser)
	appv1.Put("/users/:id", userHandler.HandleUpdateUser)

	appv1.Get("/hotels", hotelHandler.HandleGetHotels)
	appv1.Get("/hotels/:id/rooms", hotelHandler.HandleGetHotelRooms)

	app.Listen(*listenAddr)
}

package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/api"
	"github.com/seyedmahdii/hotel-reservation/api/middleware"
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

	var (
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			UserStore:    userStore,
			HotelStore:   hotelStore,
			RoomStore:    roomStore,
			BookingStore: bookingStore,
		}

		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		authHandler  = api.NewAuthHandler(userStore)
		roomHandler  = api.NewRoomHandler(store)
	)

	app := fiber.New(config)
	auth := app.Group("/api")
	appv1 := app.Group("/api/v1", middleware.JWTAuthentication(userStore))

	auth.Post("/auth", authHandler.HandleAuthenticate)

	appv1.Get("/users", userHandler.HandleGetUsers)
	appv1.Get("/users/:id", userHandler.HandleGetUser)
	appv1.Post("/users", userHandler.HandlePostUser)
	appv1.Delete("/users/:id", userHandler.HandleDeleteUser)
	appv1.Put("/users/:id", userHandler.HandleUpdateUser)

	appv1.Get("/hotels", hotelHandler.HandleGetHotels)
	appv1.Get("/hotels/:id", hotelHandler.HandleGetHotel)
	appv1.Get("/hotels/:id/rooms", hotelHandler.HandleGetHotelRooms)

	appv1.Post("/rooms/:id/book", roomHandler.HandleBookRoom)

	app.Listen(*listenAddr)
}

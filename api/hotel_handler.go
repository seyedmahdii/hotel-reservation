package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.store.HotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	hotelId := c.Params("id")
	oid, err := bson.ObjectIDFromHex(hotelId)
	if err != nil {
		return err
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.RoomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

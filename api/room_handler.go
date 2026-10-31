package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/db"
	"github.com/seyedmahdii/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

type BookRoomParams struct {
	FromDate  time.Time `json:"fromDate"`
	TillDate  time.Time `json:"tillDate"`
	NumPeople int       `json:"numPeople"`
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	roomID, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	booking := types.Booking{
		UserID:    user.ID,
		RoomID:    roomID,
		FromDate:  params.FromDate,
		TillDate:  params.TillDate,
		NumPeople: params.NumPeople,
	}

	insertedBooking, err := h.store.BookingStore.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.JSON(insertedBooking)
}

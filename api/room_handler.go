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

func (p BookRoomParams) validate() error {
	now := time.Now()

	if now.After(p.FromDate) || now.After(p.TillDate) {
		return fiber.NewError(fiber.StatusBadRequest, "dates must be in the future")
	}

	if p.TillDate.Before(p.FromDate) || p.TillDate.Equal(p.FromDate) {
		return fiber.NewError(fiber.StatusBadRequest, "tillDate must be after fromDate")
	}

	if p.NumPeople <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "numPeople must be greater than 0")
	}

	return nil
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.validate(); err != nil {
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

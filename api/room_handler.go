package api

import (
	"context"
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
		return fiber.NewError(
			fiber.StatusBadRequest,
			"dates must be in the future",
		)
	}

	if p.TillDate.Before(p.FromDate) || p.TillDate.Equal(p.FromDate) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"tillDate must be after fromDate",
		)
	}

	if p.NumPeople <= 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"numPeople must be greater than 0",
		)
	}

	return nil
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  err.Error(),
		})
	}

	if err := params.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  err.Error(),
		})
	}

	roomID, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  "invalid room ID",
		})
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	available, err := h.isRoomAvailableForBooking(c.Context(), roomID, params)
	if err != nil {
		return err
	}

	if !available {
		return c.Status(fiber.StatusConflict).JSON(genericResp{
			Type: "error",
			Msg:  "room is not available for the selected dates",
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

func (h *RoomHandler) isRoomAvailableForBooking(
	ctx context.Context,
	roomID bson.ObjectID,
	params BookRoomParams,
) (bool, error) {
	// Find existing bookings whose date ranges overlap.
	where := bson.M{
		"roomID": roomID,
		"fromDate": bson.M{
			"$lt": params.TillDate,
		},
		"tillDate": bson.M{
			"$gt": params.FromDate,
		},
	}

	bookings, err := h.store.BookingStore.GetBookings(ctx, where)
	if err != nil {
		return false, err
	}

	return len(bookings) == 0, nil
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms, err := h.store.RoomStore.GetRooms(c.Context(), bson.M{})
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

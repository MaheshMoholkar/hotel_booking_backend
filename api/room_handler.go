package api

import (
	"fmt"
	"time"

	"github.com/MaheshMoholkar/hotel_booking_backend/db"
	"github.com/MaheshMoholkar/hotel_booking_backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`
	NoOfPersons int       `json:"noOfPersons"`
}

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomOID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["id"].(string)
	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	booking := types.Booking{
		UserID:      userOID,
		RoomID:      roomOID,
		FromDate:    params.FromDate,
		ToDate:      params.ToDate,
		NoOfPersons: params.NoOfPersons,
	}
	fmt.Printf("%v", booking)

	return nil
}

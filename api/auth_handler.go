package api

import (
	"errors"

	"github.com/MaheshMoholkar/hotel_booking_backend/api/middleware"
	"github.com/MaheshMoholkar/hotel_booking_backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{userStore: userStore}
}

func (h *AuthHandler) HandleGetToken(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	id, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"msg": "Not Found"})
		}
		return err
	}
	token, err := middleware.GenerateJWT(id)
	if err != nil {
		return err
	}

	return c.JSON(token)
}

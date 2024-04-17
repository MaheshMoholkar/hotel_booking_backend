package api

import (
  "github.com/gofiber/fiber/v2"
  "github.com/MaheshMoholkar/hotel_booking_backend/types"
)

func HandleGetUser(c *fiber.Ctx) error {
  u := types.User{
    FirstName: "John",
    LastName: "Doe",
  }
  return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
  return c.JSON("John")
}

package main
import (
  "github.com/gofiber/fiber/v2"
  "flag"
  "github.com/MaheshMoholkar/hotel_booking_backend/api"
)
func main(){
  listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
  flag.Parse()

  app:=fiber.New() 
  apiv1 := app.Group("/api/v1")

  apiv1.Get("/users", api.HandleGetUsers)
  apiv1.Get("/user/:id", api.HandleGetUser)
  app.Listen(*listenAddr)

}


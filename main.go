package main

import (
	"context"
	"flag"
	"log"

	"github.com/MaheshMoholkar/hotel_booking_backend/api"
	"github.com/MaheshMoholkar/hotel_booking_backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-db"
const userColl = "users"

// var config = fiber.Config{
// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
// 		code := fiber.StatusInternalServerError

// 		var e *fiber.Error
// 		if errors.As(err, &e) {
// 			code = e.Code
// 		}

// 	},
// }

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	// handlers initializaiton
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)

}

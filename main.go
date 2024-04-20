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

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	// handlers initializaiton
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	app.Listen(*listenAddr)
}

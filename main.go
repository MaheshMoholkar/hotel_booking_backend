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

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// initializaiton
	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")
		// initialize stores
		userStore  = db.NewMongoUserStore(client, db.DBNAME)
		roomStore  = db.NewMongoRoomStore(client, db.DBNAME)
		hotelStore = db.NewMongoHotelStore(client, db.DBNAME)
		// initialize handlers
		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(hotelStore, roomStore)
	)
	// user handlers
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)

	// hotels handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/rooms/:id", hotelHandler.HandleGetRooms)
	app.Listen(*listenAddr)
}

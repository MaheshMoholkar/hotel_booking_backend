package db

const (
	DBNAME = "hotel-db"
	DBURI  = "mongodb://localhost:27017"
)

type Store struct {
	UserStore
	RoomStore
	HotelStore
}

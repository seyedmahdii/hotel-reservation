package db

const DBNAME = "hotel-reservation"
const TESTDBNAME = "hotel-reservation-test"
const DBURI = "mongodb://localhost:27017"

type Store struct {
	UserStore    UserStore
	HotelStore   HotelStore
	RoomStore    RoomStore
	BookingStore BookingStore
}

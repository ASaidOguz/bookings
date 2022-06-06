package models

import "time"

// Users is the user modal
type User struct {
	ID         int
	Firstname  string
	LastName   string
	Email      string
	Password   string
	Acceslevel int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Restrictions is the restrictions model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Reservations is the  reservations model
type Reservation struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

//RoomRestrictions is RoomRestrictions model.
type RoomRestriction struct {
	ID            int
	RoomID        int
	ReservationID int
	RestrictionID int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Reservation
}

package repository

import (
	"time"

	"github.com/ASaidOguz/bookings/internal/models"
)

type DatabaseRepo interface {
	Insertreservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilitybyDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}

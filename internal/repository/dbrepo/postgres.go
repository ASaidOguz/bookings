package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/ASaidOguz/bookings/internal/models"
)

//Insertreservation inserts reservation into database .
func (m *postgresDBRepo) Insertreservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var newID int
	sttmnt := `insert into reservations (first_name,last_name,email,phone,start_date,
	end_date,room_id,created_at,updated_at)
	values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	err := m.DB.QueryRowContext(ctx, sttmnt,
		res.Firstname,
		res.Lastname,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

//InsertRoomRestriction inserts roomrestriction into database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	sttmt := `insert into room_restrictions(start_date,end_date,room_id,reservation_id,
		    created_at,updated_at,restriction_id)
			values 
			($1,$2,$3,$4,$5,$6,$7)`

	_, err := m.DB.ExecContext(ctx, sttmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID)
	if err != nil {
		return err
	}
	return nil

}

//SearchAvailabilitybyDates true if the room is open for bussiness for roomID  false its already booked .
func (m *postgresDBRepo) SearchAvailabilitybyDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select 
              count(id)
        from
              room_restrictions
        where
		    room_id = $1
			and  
		    $2<=end_date and $3<=start_date
  `
	var numRows int

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)

	err := row.Scan(&numRows)

	if err != nil {
		//This way i find out the problem that roomID wasnt exist in my data base !!
		fmt.Println(err)
		return false, err

	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil

}

//SearchAvailabilityForAllRooms returns a slice of rooms ,if any or given date
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var rooms []models.Room

	query := `
		select
			r.id, r.room_name
		from
			rooms r
		where r.id not in 
		(select room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date);
		`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

//GetRoomByID gets room by id
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `
	select id,room_name,created_at,updated_at from rooms where id=$1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return room, err
	}

	return room, nil

}

package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//DB holds the database connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdelDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates Database pool for postgres
func ConnectSQL(dns string) (*DB, error) {
	d, err := NewDatabase(dns)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdelDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d
	err = TestDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

//TestDB tries to test ğing the database for any malfunction .
func TestDB(d *sql.DB) error {

	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

//NewDatabase creates new database for the application
func NewDatabase(dns string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dns)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

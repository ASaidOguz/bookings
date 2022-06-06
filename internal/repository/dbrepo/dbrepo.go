package dbrepo

import (
	"database/sql"

	"github.com/ASaidOguz/bookings/internal/config"
	"github.com/ASaidOguz/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func Newpostgresrepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

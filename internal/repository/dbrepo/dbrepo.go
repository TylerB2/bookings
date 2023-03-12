package dbrepo

import (
	"bookings/internal/config"
	"bookings/internal/repository"
	"database/sql"
)

type postgresDBRepo struct {
	App *config.AppConfig
	Db  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		Db:  conn,
	}

}

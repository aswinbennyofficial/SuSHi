package database

import (
	"database/sql"

	"github.com/aswinbennyofficial/SuSHi/models"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

func MigrateDB(config models.Config) error{
	log.Debug().Msg("Migrating the database")
	log.Debug().Msg("Connection string : "+ config.DatabaseConfig.String)
	db,error:=sql.Open("postgres", config.DatabaseConfig.String)
	if error != nil {
		return error
	}
	defer db.Close()

	goose.SetDialect("postgres")

	err:=goose.Up(db, "db/migrations")
	if err != nil {
		return err
	}
	
	return nil
}
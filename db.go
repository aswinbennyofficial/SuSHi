package main

import (
	"context"
	"time"


	"github.com/jackc/pgx/v5/pgxpool"
	
	"github.com/pressly/goose/v3"
	"database/sql"
	_ "github.com/lib/pq"
	
	"github.com/rs/zerolog/log"

	
)

// ConnectDB() is used to connect to the database using the configuration values. It initializes the package-level variable DB with the database connection pool. It returns the error as output.
func (config *Config)ConnectDB() error {
	connectionString := "host=" + config.DatabaseConfig.Host + 
		" port=" + config.DatabaseConfig.Port + 
		" user=" + config.DatabaseConfig.User + 
		" password=" + config.DatabaseConfig.Password + 
		" dbname=" + config.DatabaseConfig.Database + 
		" sslmode=disable"

	config.DatabaseConfig.String = connectionString

	ctx := context.Background()
	var conn *pgxpool.Pool
	var err error

	for i := 0; i < 5; i++ {
		conn, err = pgxpool.New(ctx, connectionString)
		if err == nil {
			break
		}
		log.Error().Msgf("Attempt %d: Failed to connect to the database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal().Msgf("Failed to connect to the database after 3 attempts: %v", err)
		return err
	}

	

	// ping the database
	for i := 0; i < 5; i++ {
		err = conn.Ping(ctx)
		if err == nil {
			break
		}
		log.Error().Msgf("Attempt %d: Failed to ping the database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	log.Info().Msg("Connected to the database")

	config.DB = conn
	return nil
}


func (config *Config)migrateDB() error{
	
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
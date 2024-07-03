package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

// DB is the package-level variable which is used to store the database connection pool
var DB *pgx.Conn

// ConnectDB() is used to connect to the database using the configuration values. It initializes the package-level variable DB with the database connection pool. It returns the error as output.
func ConnectDB() error {
	connectionString := "host=" + config.DatabaseConfig.Host + 
		" port=" + config.DatabaseConfig.Port + 
		" user=" + config.DatabaseConfig.User + 
		" password=" + config.DatabaseConfig.Password + 
		" dbname=" + config.DatabaseConfig.Database + 
		" sslmode=disable"

	ctx := context.Background()
	var conn *pgx.Conn
	var err error

	for i := 0; i < 3; i++ {
		conn, err = pgx.Connect(ctx, connectionString)
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

	log.Info().Msg("Connected to the database")

	// ping the database
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal().Msgf("Failed to ping the database: %v", err)
		return err
	}

	DB = conn
	return nil
}

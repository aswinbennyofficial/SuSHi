package database

import (
	"context"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/rs/zerolog/log"
)

func DoesUserExists(config models.Config, email string)(bool){

	var count int
	err := config.DB.QueryRow(context.Background(), "SELECT COUNT(email) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		// Handle any errors during the query execution
		log.Debug().Msgf("Error checking if user exists: %v", err)
		return false
	}
	log.Debug().Msgf("DoesUserExists() : Count: %d",count)
	return count > 0

}

func SaveUser(config models.Config, email string, name string) (error) {

	salt,err := utils.GenerateUUID(8)
	if err != nil {
		log.Error().Msgf("Error generating salt: %v", err)
		return err
	}

	// save user to database
	_, err = config.DB.Exec(context.Background(), "INSERT INTO users (email, name, salt) VALUES ($1, $2, $3)", email, name, salt)
	if err != nil {
		log.Error().Msgf("Error saving user to database: %v", err)
		return err
	}
	return nil
}

func GetUsername(config models.Config, email string)(string,error){

	var username string
	err := config.DB.QueryRow(context.Background(), "SELECT username FROM users WHERE email = $1", email).Scan(&username)
	if err != nil {
		// Handle any errors during the query execution
		log.Debug().Msgf("Error fetching username: %v", err)
		return "",err
	}
	log.Debug().Msgf("GetUsername() : Username: %s",username)
	return username,nil

}
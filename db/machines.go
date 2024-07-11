package database

import (
	"context"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/rs/zerolog/log"
)


func SaveMachine(config models.Config, machine models.Machine, username string) (error) {
	// fetch salt from database
	var salt string
	err := config.DB.QueryRow(context.Background(), "SELECT salt FROM users WHERE username = $1", username).Scan(&salt)
	if err != nil {
		log.Error().Msgf("Error fetching salt from database: %v", err)
		return err
	}
	
	log.Debug().Msg("SaveMachine() : Salt: "+salt)


	enc_text,ivString,err:=utils.EncryptString(machine.PrivateKey, machine.Password, salt)
	if err != nil {
		log.Error().Msgf("Error encrypting private key: %v", err)
		return err
	}


	owner_id:=username
	owner_type:="user"


	// owner type 
	if machine.Organization != ""{
		owner_id=machine.Organization
		owner_type="organization"
	}


	// save machine to database
	_, err = config.DB.Exec(context.Background(), "INSERT INTO machines (name, username, hostname, port, encrypted_private_key, encrypted_passphrase, iv, owner_id, owner_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", machine.Name, machine.Username, machine.Hostname, machine.Port, enc_text, ivString, salt, owner_id, owner_type)
	if err != nil {
		log.Error().Msgf("Error saving machine to database: %v", err)
		return err
	}
	log.Debug().Msg("Machine saved to database")

	return nil

}
package database

import (
	"context"
	"errors"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)


func SaveMachine(config models.Config, machine models.MachineRequest, username string) (error) {
	// fetch salt from database
	var salt string
	err := config.DB.QueryRow(context.Background(), "SELECT salt FROM users WHERE username = $1", username).Scan(&salt)
	if err != nil {
		log.Error().Msgf("Error fetching salt from database: %v", err)
		return err
	}
	
	log.Debug().Msg("SaveMachine() : Salt: "+salt)


	enc_PrivateKey,ivPrivateKey,err:=utils.EncryptString(machine.PrivateKey, machine.Password, salt)
	if err != nil {
		log.Error().Msgf("Error encrypting private key: %v", err)
		return err
	}

	enc_Passphrase:=""
	ivPassphrase:=""
	if machine.Passphrase!=""{
		enc_Passphrase, ivPassphrase, err = utils.EncryptString(machine.Passphrase, machine.Password, salt)
		if err != nil {
			log.Error().Msgf("Error encrypting passphrase: %v", err)
			return err
		}
	}
	


	owner_id:=username
	owner_type:="user"


	// owner type 
	if machine.Organization != ""{
		owner_id=machine.Organization
		owner_type="organization"
	}


	// save machine to database
	_, err = config.DB.Exec(context.Background(), "INSERT INTO machines (name, username, hostname, port, encrypted_private_key, iv_private_key, encrypted_passphrase, iv_passphrase, owner_id, owner_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", machine.Name, machine.Username, machine.Hostname, machine.Port, enc_PrivateKey, ivPrivateKey, enc_Passphrase, ivPassphrase, owner_id, owner_type)
	if err != nil {
		log.Error().Msgf("Error saving machine to database: %v", err)
		return err
	}
	log.Debug().Msg("Machine saved to database")

	return nil

}


func GetMachines(config models.Config, user_id string, owner_type string) ([]models.Machine, error) {
	var machines []models.Machine

	rows, err := config.DB.Query(context.Background(), "SELECT id, name, username, hostname, port, encrypted_private_key, iv_private_key, encrypted_passphrase, iv_passphrase FROM machines WHERE owner_id = $1 AND owner_type = $2", user_id, owner_type)

	if err != nil {
		log.Error().Msgf("Error fetching machines from database: %v", err)
		return machines, err
	}
	defer rows.Close()

	for rows.Next() {
		var machine models.Machine
		err := rows.Scan(&machine.ID, &machine.Name, &machine.Username, &machine.Hostname, &machine.Port, &machine.PrivateKey, &machine.IvPrivateKey, &machine.Passphrase, &machine.IvPassphrase)
		if err != nil {
			log.Error().Msgf("Error scanning machine from database: %v", err)
			return machines, err
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func GetMachinesBasicInfo(config models.Config, user_id string, owner_type string) ([]models.FilterMachine, error) {
	var machines []models.FilterMachine

	rows, err := config.DB.Query(context.Background(), "SELECT id, name, username, hostname, port, owner_id, owner_type FROM machines WHERE owner_id = $1 AND owner_type = $2", user_id, owner_type)

	if err != nil {
		log.Error().Msgf("Error fetching machines from database: %v", err)
		return machines, err
	}
	defer rows.Close()

	for rows.Next() {
		var machine models.FilterMachine
		err := rows.Scan(&machine.ID, &machine.Name, &machine.Username, &machine.Hostname, &machine.Port, &machine.OwnerID, &machine.OwnerType)
		if err != nil {
			log.Error().Msgf("Error scanning machine from database: %v", err)
			return machines, err
		}
		machines = append(machines, machine)
	}

	return machines, nil
}


func GetAMachineBasicInfo(config models.Config, machine_id string, user_id string, owner_type string) (models.FilterMachine, error) {
	var machine models.FilterMachine

	err := config.DB.QueryRow(context.Background(), "SELECT id, name, username, hostname, port, owner_id, owner_type FROM machines WHERE id = $1 AND owner_id = $2 AND owner_type = $3", machine_id,user_id,owner_type).Scan(&machine.ID, &machine.Name, &machine.Username, &machine.Hostname, &machine.Port, &machine.OwnerID, &machine.OwnerType)

	

	if err != nil {

		if err == pgx.ErrNoRows {
			log.Error().Msgf("Machine not found in database")
			return machine, errors.New("machine not found in database")
		}

		log.Error().Msgf("Error fetching machine from database: %v", err)
		return machine, err
	}

	return machine, nil
}

func GetAMachine(config models.Config, machine_id string, user_id string, owner_type string, password string) (models.Machine, error) {
	var machine models.Machine

	err := config.DB.QueryRow(context.Background(), "SELECT id, name, username, hostname, port, encrypted_private_key, iv_private_key, encrypted_passphrase, iv_passphrase FROM machines WHERE id = $1 AND owner_id = $2 AND owner_type = $3", machine_id,user_id,owner_type).Scan(&machine.ID, &machine.Name, &machine.Username, &machine.Hostname, &machine.Port, &machine.PrivateKey, &machine.IvPrivateKey, &machine.Passphrase, &machine.IvPassphrase)

	if err != nil {

		if err == pgx.ErrNoRows {
			log.Error().Msgf("Machine not found in database")
			return machine, errors.New("machine not found in database")
		}

		log.Error().Msgf("Error fetching machine from database: %v", err)
		return machine, err
	}

	// decrypt private key
	// fetch salt from database
	var salt string
	err = config.DB.QueryRow(context.Background(), "SELECT salt FROM users WHERE username = $1", user_id).Scan(&salt)
	if err != nil {
		log.Error().Msgf("Error fetching salt from database: %v", err)
		return machine, err
	}

	machine.PrivateKey,err=utils.DecryptString(machine.PrivateKey, machine.IvPrivateKey, password, salt)
	if err != nil {
		log.Error().Msgf("Error decrypting private key: %v", err)
		return machine, err
	}

	if machine.Passphrase!=""{
		log.Debug().Msg("Passphrase:"+machine.Passphrase)
		machine.Passphrase,err=utils.DecryptString(machine.Passphrase, machine.IvPassphrase, password, salt)
		if err != nil {
			log.Error().Msgf("Error decrypting passphrase: %v", err)
			return machine, err
		}
	}

	return machine, nil
}


func DeleteMachine(config models.Config, machine_id string, user_id string, owner_type string) error {
	_, err := config.DB.Exec(context.Background(), "DELETE FROM machines WHERE id = $1 AND owner_id = $2 AND owner_type = $3", machine_id, user_id, owner_type)
	if err != nil {
		log.Error().Msgf("Error deleting machine from database: %v", err)
		return err
	}
	log.Debug().Msg("Machine deleted from database")
	return nil
}


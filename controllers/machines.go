package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/aswinbennyofficial/SuSHi/db"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/rs/zerolog/log"
	"github.com/go-chi/chi/v5"
)

func CreateMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	var machine models.Machine
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&machine)

	// debug log
	log.Debug().Msg("Machine name: "+machine.Name)
	log.Debug().Msg("Machine password: "+machine.Password)
	log.Debug().Msg("Machine username: "+machine.Username)
	log.Debug().Msg("Machine hostname: "+machine.Hostname)
	log.Debug().Msgf("Machine port:"+machine.Port)
	log.Debug().Msg("Machine private key: "+machine.PrivateKey)
	log.Debug().Msg("Machine passphrase: "+machine.Passphrase)
	log.Debug().Msg("Machine organization: "+machine.Organization)




	// fetch username from jwt token
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching username from token", err)
		return
	}
	log.Debug().Msg("Username: "+username)

	// save machine to database
	err=database.SaveMachine(config, machine, username)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error saving machine to database", err)
		return
	}
	

	// test response
	utils.ResponseHelper(w, http.StatusOK, "Machine created successfully", nil)

}



func GetMachines(config models.Config, w http.ResponseWriter, r *http.Request){
	
	
	// fetch username from jwt token
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching username from token", err)
		return
	}
	log.Debug().Msg("Username: "+username)

	// fetch machines from database
	machines,err:=database.GetMachinesBasicInfo(config, username,"user")
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching machines from database", err)
		return
	}

	

	


	utils.ResponseHelper(w, http.StatusOK, "Machines fetched successfully", machines)
}

func GetMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	
	// fetch username from jwt token
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching username from token", err)
		return
	}
	log.Debug().Msg("Username: "+username)

	// fetch machine id from url params
	machineID := chi.URLParam(r, "id")

	log.Debug().Msg("Machine ID: "+machineID)

	// fetch machine from database
	machine,err:=database.GetAMachineBasicInfo(config, machineID, username, "user")
	if err != nil {

		if err.Error()=="machine not found in database"{
			utils.ResponseHelper(w, http.StatusNotFound, "Machine not found in database", nil)
			return
		}

		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching machine from database", err)
		return
	}
	
	utils.ResponseHelper(w, http.StatusOK, "Machine fetched successfully", machine)
}



func DeleteMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	// fetch username from jwt token
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching username from token", err)
		return
	}
	log.Debug().Msg("Username: "+username)

	// fetch machine id from url params
	machineID := chi.URLParam(r, "id")

	log.Debug().Msg("Machine ID: "+machineID)

	// delete machine from database
	err=database.DeleteMachine(config, machineID, username, "user")
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error deleting machine from database", err)
		return
	}


	utils.ResponseHelper(w, http.StatusOK, "Machine deleted successfully", nil)
}
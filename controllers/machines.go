package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/aswinbennyofficial/SuSHi/db"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func CreateMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	var machine models.MachineRequest
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


func ConnectMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching username from token", err)
		return
	}
	log.Debug().Msg("Username: "+username)

	// fetch machine id from url params
	machineID := chi.URLParam(r, "id")

	log.Debug().Msg("Machine ID: "+machineID)

	// get password from post request
	var requestBody models.ConnectionRequest

	err=json.NewDecoder(r.Body).Decode(&requestBody)
	if err !=nil{
		log.Debug().Msg("Error decoding post body"+err.Error())
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error decoding post body", err)
	}

	log.Debug().Msg("Password: "+requestBody.Password)
	


	machine,err:=database.GetAMachine(config,machineID,username,"user",requestBody.Password)
	if err != nil {
		if err.Error()=="machine not found in database"{
			utils.ResponseHelper(w, http.StatusNotFound, "Machine not found in database", nil)
			return
		}

		utils.ResponseHelper(w, http.StatusInternalServerError, "Error fetching machine from database", err)
		return
	}

	// use ssh to connect to machine
	sshClient,err:=utils.ConnectToMachine(machine)
	if err != nil {
		utils.ResponseHelper(w, http.StatusInternalServerError, "Error connecting to machine", err)
		return
	}
	log.Debug().Msg("Connected to machine successfully")

	// generate uuid
	uuid := uuid.New().String()
	log.Debug().Msg("UUID: "+uuid)

	// store ssh connection
	timeKey:=utils.RoundToNearestMinute(time.Now())
	utils.StoreSSHConnection(uuid, &models.SSHConnection{TimeBucketKey: timeKey, Client: sshClient})
	utils.StoreInTimeBucket(timeKey, uuid)
	log.Debug().Msg("SSH connection stored successfully"+uuid+" "+timeKey.String())

	// send response
	utils.ResponseHelper(w, http.StatusOK, "Connected to machine successfully", uuid)


	// utils.ResponseHelper(w, http.StatusOK, "Machine fetched successfully", machine)
}
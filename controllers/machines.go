package controllers

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/rs/zerolog/log"
)

func CreateMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	

	// fetch username from jwt token
	username,err:=utils.GetUsernameFromToken(r)
	if err != nil {
		utils.ResponseHelper(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	log.Debug().Msg("Username: "+username)
	

	// test response
	utils.ResponseHelper(w, http.StatusOK, "Machine created successfully", nil)

}



func GetMachines(config models.Config, w http.ResponseWriter, r *http.Request){
	// test response
	utils.ResponseHelper(w, http.StatusOK, "Machines fetched successfully", nil)
}

func GetMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	
	// test response
	utils.ResponseHelper(w, http.StatusOK, "Machine fetched successfully", nil)
}

func DeleteMachine(config models.Config, w http.ResponseWriter, r *http.Request){
	utils.ResponseHelper(w, http.StatusOK, "Machine deleted successfully", nil)
}
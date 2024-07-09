package controllers

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"
)

func CreateMachine(config models.Config, w http.ResponseWriter, r *http.Request){
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
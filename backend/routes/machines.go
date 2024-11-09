package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/controllers"
	"github.com/aswinbennyofficial/SuSHi/models"

	"github.com/go-chi/chi/v5"
	
)


func loadMachineRoutes(r chi.Router, config models.Config){

	r.Post("/machine", func(w http.ResponseWriter, r *http.Request) {
        controllers.CreateMachine(config, w, r)
    })

	r.Get("/machines", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetMachines(config, w, r)
		
	})

	r.Get("/machine/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetMachine(config, w, r)
		
	})

	r.Delete("/machine/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteMachine(config, w, r)
		
	})

	r.Post("/machine/{id}/connect", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectMachine(config, w, r)
		
	})

	


}
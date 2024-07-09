package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/controllers"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/go-chi/chi/v5"
)


func loadSSHRoutes(r chi.Router, config models.Config){
	// WebSocket handler
    r.Get("/ssh", func(w http.ResponseWriter, r *http.Request) {
        controllers.HandleSSHConnection(config, w, r)
		
    })
}
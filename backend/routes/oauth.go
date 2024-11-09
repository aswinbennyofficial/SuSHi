package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/controllers"
	"github.com/aswinbennyofficial/SuSHi/models"

	"github.com/go-chi/chi/v5"
)

func loadoAuthRoutes(r chi.Router, config models.Config){
	r.Get("/url", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAuthURL(config,w,r)
	})

	r.Get("/callback",func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleCallback(config,w,r)
	})

	r.Get("/logout",func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleLogout(config,w,r)
	})
}
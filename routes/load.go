package routes

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/rs/zerolog/log"
)

func Load(config models.Config) {
	r:=config.Router
	// Serve static files (e.g., xterm.js frontend)
	workDir, _ := os.Getwd()
	filesDir := http.Dir(workDir + "/static")
    log.Print(workDir+"/static")
	fileServer := http.FileServer(filesDir)
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// JWT authentication middleware
	tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
        r.Use(jwtauth.Authenticator(tokenAuth))

		// Define API version base path
		r.Route("/api/v1", func(r chi.Router) {
			loadMachineRoutes(r, config)
		})
	})

	// Load routes
	loadSSHRoutes(r, config)
	



	
	
}
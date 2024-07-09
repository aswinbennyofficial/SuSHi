package routes

import (
	"net/http"
	"os"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/ssh"
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

	
	// WebSocket handler
    r.Get("/ssh", func(w http.ResponseWriter, r *http.Request) {
        ssh.HandleSSHConnection(config, w, r)
    })
}
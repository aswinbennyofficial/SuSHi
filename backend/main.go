package main

import (
	
	"net/http"
	"os"
	

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
)


func init(){
	//  Load configuration
    err := LoadConfig()
    if err != nil {
        log.Panic().Err(err).Msg("Error in LoadConfig()")
        return
    }

    log.Info().Msg("Configuration loaded successfully")

    // Set log level
    if config.LogLevel == "Debug" {
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
    } else {
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
    }

    // Load logger
    LoadLogger()



	// Show config log
	showConfigLog()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func main() {
	r := chi.NewRouter()

	// Serve static files (e.g., xterm.js frontend)
	workDir, _ := os.Getwd()
	filesDir := http.Dir(workDir + "/static")
    log.Print(workDir+"/static")
	fileServer := http.FileServer(filesDir)
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// WebSocket handler
	r.Get("/ssh", handleSSHConnection)

	// Start HTTP server
	log.Info().Msgf("Starting server on port %s", config.ServerPort)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal().Msgf("Failed to start server: %v", err)
	}
}


func showConfigLog(){
	log.Debug().Msgf("ServerPort: %s", config.ServerPort)
	log.Debug().Msgf("JWTSecret: %s", config.JWTSecret)
	log.Debug().Msgf("LogLevel: %s", config.LogLevel)
	log.Debug().Msgf("LogPath: %s", config.LogPath)
	log.Debug().Msgf("SSHHost: %s", config.SSHConfig.SSHHost)
	log.Debug().Msgf("SSHPort: %s", config.SSHConfig.SSHPort)
	log.Debug().Msgf("SSHUser: %s", config.SSHConfig.SSHUser)
	log.Debug().Msgf("PrivateKey: %s", config.SSHConfig.PrivateKey)
	log.Debug().Msgf("Host: %s", config.DatabaseConfig.Host)
	log.Debug().Msgf("Port: %s", config.DatabaseConfig.Port)
	log.Debug().Msgf("User: %s", config.DatabaseConfig.User)
	log.Debug().Msgf("Password: %s", config.DatabaseConfig.Password)
	log.Debug().Msgf("Database: %s", config.DatabaseConfig.Database)

}
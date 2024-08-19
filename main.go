package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	database "github.com/aswinbennyofficial/SuSHi/db"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/routes"
	"github.com/aswinbennyofficial/SuSHi/utils"
)







func main() {
	var config models.Config

	//  Load configuration
    config,err := utils.LoadConfig()
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
    utils.LoadLogger(config)



	// Show config log
	showConfigLog(config)

	

	// Connect to the database
	config.DB,config.DatabaseConfig.String,err = database.Connect(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
		return
	}


	// Do migration
	err = database.Migrate(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate the database")
		return
	}

	// Start goroutine to remove expired ssh connections
	go utils.CheckExpiredBuckets()
	
	

	config.Router = chi.NewRouter()

	
	routes.Load(config)

	// Start HTTP server
	log.Info().Msgf("Starting server on port %s", config.ServerPort)
	err = http.ListenAndServe(":8080", config.Router)
	if err != nil {
		log.Fatal().Msgf("Failed to start server: %v", err)
	}
}


func showConfigLog(config models.Config){
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
	log.Debug().Msgf("Google : cid %s | csecret %s | scope %s | redirecturl %s | state %s", config.OAuthConfig.Google.ClientID,config.OAuthConfig.Google.ClientSecret,config.OAuthConfig.Google.Scopes,config.OAuthConfig.Google.RedirectURL,config.OAuthConfig.Google.State)
	log.Debug().Msgf("Github client id : %s", config.OAuthConfig.GitHub.ClientID)

}


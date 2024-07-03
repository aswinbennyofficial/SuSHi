package main

import (
	"os"
	// "strconv"
	"strings"


    // "github.com/rs/zerolog/log"
)

// config is the package-level variable which is used to store the configuration
var config Config

/*
LoadConfig() is used to load the configuration from the environment variables . It initializes the package-level variable config  with the configuration values. It returns the error as output.
*/
func LoadConfig() error {
    
    config.ServerPort = os.Getenv("SERVER_PORT")
    config.JWTSecret = os.Getenv("JWT_SECRET")
	config.LogLevel = os.Getenv("LOG_LEVEL")
    config.LogPath = os.Getenv("LOG_PATH")

    // SSH Config
    config.SSHConfig.SSHHost = os.Getenv("SSH_HOST")
    config.SSHConfig.SSHPort = os.Getenv("SSH_PORT")
    config.SSHConfig.SSHUser = os.Getenv("SSH_USER")
    config.SSHConfig.PrivateKey = os.Getenv("SSH_PRIVATE_KEY")

    // Database Config
    config.DatabaseConfig.Host = os.Getenv("DB_HOST")
    config.DatabaseConfig.Port = os.Getenv("DB_PORT")
    config.DatabaseConfig.User = os.Getenv("DB_USER")
    config.DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
    config.DatabaseConfig.Database = os.Getenv("DB_NAME")


   

    // Use default values if environment variables are not set
    if config.LogPath == "" {
        config.LogPath = "./logs"
    } else {
        // remove any trailing slashes
        config.LogPath = strings.TrimSuffix(config.LogPath, "/")
    }


    return nil
}
package utils

import (
	"fmt"
	"os"
	// "strconv"
	"strings"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/spf13/viper"
	// "github.com/rs/zerolog/log"
)

/*
LoadConfig() is used to load the configuration from the environment variables . It initializes the package-level variable config  with the configuration values. It returns the error as output.
*/
func LoadConfig() (models.Config, error) {
    var config models.Config

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

    config.DoMigrations = (os.Getenv("MIGRATE_DB") == "true")


    // Use default values if environment variables are not set
    if config.LogPath == "" {
        config.LogPath = "./logs"
    } else {
        // remove any trailing slashes
        config.LogPath = strings.TrimSuffix(config.LogPath, "/")
    }


    viper.SetConfigName("oauth")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")
    viper.SetEnvPrefix("") // This allows Viper to see env vars without a prefix
    viper.AutomaticEnv()   // Tells Viper to check for env vars

    if err := viper.ReadInConfig(); err != nil {
        return config, fmt.Errorf("error reading config file: %w", err)
    }

    var oauthConfig models.OAuthConfig
    if err := viper.Sub("oauth").Unmarshal(&oauthConfig); err != nil {
        return config, fmt.Errorf("error unmarshalling config: %w", err)
    }

    // Expand environment variables
    oauthConfig.Google.ClientID = viper.GetString("GOOGLE_CLIENT_ID")
    oauthConfig.Google.ClientSecret = viper.GetString("GOOGLE_CLIENT_SECRET")
    oauthConfig.Google.RedirectURL = viper.GetString("GOOGLE_REDIRECT_URL")
    oauthConfig.GitHub.ClientID = viper.GetString("GITHUB_CLIENT_ID")
    oauthConfig.GitHub.ClientSecret = viper.GetString("GITHUB_CLIENT_SECRET")
    oauthConfig.GitHub.RedirectURL = viper.GetString("GITHUB_REDIRECT_URL")

    config.OAuthConfig=oauthConfig

    return config,nil
}
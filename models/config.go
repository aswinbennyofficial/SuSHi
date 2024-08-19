package models

import(
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"
)


// Config struct
type Config struct {
	
	ServerPort         string
	JWTSecret          string
	LogLevel 		   string
	LogPath            string
	SSHConfig          SSHConfig
	DatabaseConfig     DatabaseConfig
	DB                 *pgxpool.Pool
	Router			   *chi.Mux
	DoMigrations	   bool
	OAuthConfig 		OAuthConfig

}

type SSHConfig struct {
	SSHHost string
	SSHPort string
	SSHUser string
	PrivateKey string
}

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	String 	string
}




type ProviderConfig struct {
    ClientID     string   `yaml:"client_id"`
    ClientSecret string   `yaml:"client_secret"`
    Scopes       []string `yaml:"scopes"`
    RedirectURL  string   `yaml:"redirect_url"`
    State        string   `yaml:"state"`
}

type OAuthConfig struct {
    Google ProviderConfig `yaml:"google"`
    GitHub ProviderConfig `yaml:"github"`
}

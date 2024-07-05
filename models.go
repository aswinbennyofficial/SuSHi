package main

import(
	"github.com/jackc/pgx/v5/pgxpool"
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


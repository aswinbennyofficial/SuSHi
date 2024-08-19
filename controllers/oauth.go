package controllers

import (
	"net/http"
	"time"

	database "github.com/aswinbennyofficial/SuSHi/db"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/oauth"
	"github.com/aswinbennyofficial/SuSHi/utils"
	"github.com/rs/zerolog/log"
)

func GetAuthURL(config models.Config,w http.ResponseWriter, r *http.Request){
	// fetch the query param 'to'
	vendor := r.URL.Query().Get("to")

	if vendor == "" {
		http.Error(w, "Missing query param 'to'", http.StatusBadRequest)
		return
	}

	
	var authURL string
	switch vendor {
	case "google":
		authURL = oauth.GenerateGoogleAuthURL(config)
	case "github":
		authURL = oauth.GenerateGitHubAuthURL(config)
	default:
		http.Error(w, "Invalid query param 'to'", http.StatusBadRequest)
		return
	}

	w.Write([]byte(authURL))

}


func HandleCallback(config models.Config,w http.ResponseWriter, r *http.Request){
	// check the state parameter
	state := r.URL.Query().Get("state")
	if state == "" {
		http.Error(w, "Missing query param 'state'", http.StatusBadRequest)
		return
	}

	// check the code parameter
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing query param 'code'", http.StatusBadRequest)
		return
	}

	var email string
	var name string
	var err error

	switch state {
	case config.OAuthConfig.Google.State:
		email, name, err = oauth.HandleGoogleCallback(config, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		doesUserExists:=database.DoesUserExists(config,email)
		if !doesUserExists{
			log.Debug().Msgf("User %s does not exist",email)
			err=database.SaveUser(config,email,name)
			if err != nil {
				log.Error().Msgf("Error saving user to database: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	case config.OAuthConfig.GitHub.State:
		email, name, err = oauth.HandleGitHubCallback(config, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		doesUserExists:=database.DoesUserExists(config,email)
		if !doesUserExists{
			log.Debug().Msgf("User %s does not exist",email)
			err=database.SaveUser(config,email,name)
			if err != nil {
				log.Error().Msgf("Error saving user to database: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	default:
		http.Error(w, "Invalid query param 'state'", http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("Email: %s", email)
	// Fetch username from the database
	username, err := database.GetUsername(config, email)
	if err != nil {
		http.Error(w, "Error fetching username from database", http.StatusInternalServerError)
		return
	} 
	log.Debug().Msgf("Username: %s", username)

	// Generate a JWT token
	token, exp, err := utils.GenerateJWT(config.JWTSecret, username)
	if err != nil {
		log.Error().Msgf("Error generating JWT token: %v", err)
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}
	log.Debug().Msgf("JWT token: %s", token)

	// Set the JWT token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: false, //@TODO: Change to true
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		Expires: exp,
		Path: "/",
	})

	// Redirect the user to the dashboard
	http.Redirect(w, r, "/dashboard.html", http.StatusSeeOther)

}

func HandleLogout(config models.Config,w http.ResponseWriter, r *http.Request){

	// set expiry time to 10 seconds ago
	exp := time.Now().Add(-10 * time.Second)

	// Set the JWT token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: false, //@TODO: Change to true
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		Expires: exp,
		Path: "/",
	})

	utils.ResponseHelper(w, http.StatusOK, "Logged out successfully", nil)

}
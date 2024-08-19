package controllers

import (
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/oauth"
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

	switch state {
	case config.OAuthConfig.Google.State:
		email, name, err := oauth.HandleGoogleCallback(config, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Email: " + email + "\nName: " + name))
	case config.OAuthConfig.GitHub.State:
		email, name, err := oauth.HandleGitHubCallback(config, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Email: " + email + "\nName: " + name))
	default:
		http.Error(w, "Invalid query param 'state'", http.StatusBadRequest)
		return
	}
}
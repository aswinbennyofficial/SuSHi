package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GenerateGoogleAuthURL(config models.Config) string {
	googleConfig := config.OAuthConfig.Google

    // Prepare the query parameters
    params := url.Values{}
    params.Add("client_id", googleConfig.ClientID)
    params.Add("redirect_uri", googleConfig.RedirectURL)
    params.Add("response_type", "code")
    params.Add("scope", strings.Join(googleConfig.Scopes, " "))
    params.Add("state", googleConfig.State)

    // Construct the final URL
    authURL := "https://accounts.google.com/o/oauth2/auth?" + params.Encode()
    
    return authURL
}

// HandleGoogleCallback handles the Google OAuth callback, exchanges the code for an access token, and fetches the user's email and name.
func HandleGoogleCallback(config models.Config, code string) (string, string, error) {
    googleConfig := config.OAuthConfig.Google

    // Create the OAuth2 config with the necessary credentials and redirect URL
    oauth2Config := &oauth2.Config{
        ClientID:     googleConfig.ClientID,
        ClientSecret: googleConfig.ClientSecret,
        RedirectURL:  googleConfig.RedirectURL,
        Scopes:       googleConfig.Scopes,
        Endpoint:     google.Endpoint,
    }

    // Exchange the authorization code for an access token
    token, err := oauth2Config.Exchange(context.Background(), code)
    if err != nil {
        return "", "", fmt.Errorf("failed to exchange code for token: %w", err)
    }

    // Create a new HTTP client using the access token
    client := oauth2Config.Client(context.Background(), token)

    // Make a request to the Google API to fetch the user's profile information
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json")
    if err != nil {
        return "", "", fmt.Errorf("failed to get user info: %w", err)
    }
    defer resp.Body.Close()

    // Decode the response into a struct
    var userInfo struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return "", "", fmt.Errorf("failed to decode user info: %w", err)
    }

	log.Debug().Msgf("User info: %+v", userInfo)
	log.Debug().Msgf("User email: %s", userInfo.Email)
	
    return userInfo.Email, userInfo.Name, nil
}
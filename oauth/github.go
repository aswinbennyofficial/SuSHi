package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/aswinbennyofficial/SuSHi/models"
	"golang.org/x/oauth2"
)


func GenerateGitHubAuthURL(config models.Config) string {
    githubConfig := config.OAuthConfig.GitHub

    // Prepare the query parameters
    params := url.Values{}
    params.Add("client_id", githubConfig.ClientID)
    params.Add("redirect_uri", githubConfig.RedirectURL)
    params.Add("scope", strings.Join(githubConfig.Scopes, " "))
    params.Add("state", githubConfig.State)

    // Construct the final URL
    authURL := "https://github.com/login/oauth/authorize?" + params.Encode()

    return authURL
}


// HandleGitHubCallback handles the GitHub OAuth callback, exchanges the code for an access token, and fetches the user's email and name.
func HandleGitHubCallback(config models.Config, code string) (string, string, error) {
    githubConfig := config.OAuthConfig.GitHub

    // Create the OAuth2 config with the necessary credentials and redirect URL
    oauth2Config := &oauth2.Config{
        ClientID:     githubConfig.ClientID,
        ClientSecret: githubConfig.ClientSecret,
        RedirectURL:  githubConfig.RedirectURL,
        Scopes:       githubConfig.Scopes,
        Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
    }

    // Exchange the authorization code for an access token
    token, err := oauth2Config.Exchange(context.Background(), code)
    if err != nil {
        return "", "", fmt.Errorf("failed to exchange code for token: %w", err)
    }

    // Create a new HTTP client using the access token
    client := oauth2Config.Client(context.Background(), token)

    // Make a request to the GitHub API to fetch the user's profile information
    resp, err := client.Get("https://api.github.com/user")
    if err != nil {
        return "", "", fmt.Errorf("failed to get user info: %w", err)
    }
    defer resp.Body.Close()

    // Decode the response into a struct to get the user's name
    var userInfo struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return "", "", fmt.Errorf("failed to decode user info: %w", err)
    }

    // Fetch the user's email separately if it's not provided in the main user info (GitHub might not always provide the email directly)
    if userInfo.Email == "" {
        emailResp, err := client.Get("https://api.github.com/user/emails")
        if err != nil {
            return "", "", fmt.Errorf("failed to get user emails: %w", err)
        }
        defer emailResp.Body.Close()

        var emails []struct {
            Email   string `json:"email"`
            Primary bool   `json:"primary"`
            Verified bool  `json:"verified"`
        }
        if err := json.NewDecoder(emailResp.Body).Decode(&emails); err != nil {
            return "", "", fmt.Errorf("failed to decode emails: %w", err)
        }

        // Find the primary and verified email
        for _, email := range emails {
            if email.Primary && email.Verified {
                userInfo.Email = email.Email
                break
            }
        }
    }

    // If no name is provided, fallback to the username
    if userInfo.Name == "" {
        userInfo.Name = strings.Split(userInfo.Email, "@")[0]
    }

    return userInfo.Email, userInfo.Name, nil
}

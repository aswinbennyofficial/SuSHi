package utils

import (
	"errors"
	
	"net/http"
	

	
	"time"

	"github.com/go-chi/jwtauth/v5"
	
)

func GetUsernameFromToken(r *http.Request) (string, error) {
	// fetch username from jwt token
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return "", errors.New("error fetching claims from token"+err.Error())
	}
    user_id, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("no user_id in token claims")
	}

	return user_id, nil


}


func GenerateJWT(secret string, username string) (string, time.Time, error) {
    // Create a new token auth instance with the secret
    tokenAuth := jwtauth.New("HS256", []byte(secret), nil)

    // Current time
    now := time.Now()
	exp := now.Add(24 * time.Hour)

    // Create claims for the token
    claims := map[string]interface{}{
        "username": username,
        "iat": now.Unix(),
        "exp": exp, // Token expires in 24 hours
    }

	

    // Create the JWT
    _, tokenString, err := tokenAuth.Encode(claims)
    if err != nil {
        return "", exp, err
    }

    return tokenString, exp, nil
}
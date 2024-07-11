package utils

import (
	"errors"
	"net/http"

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
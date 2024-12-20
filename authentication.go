package passage

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// AuthenticateRequest validates the JWT from either the Authorization header or cookie and returns the user ID.
//
// Deprecated: use `Passage.Auth.ValidateJWT` instead.
func (a *App) AuthenticateRequest(r *http.Request) (string, error) {
	if a.Config.HeaderAuth {
		return a.AuthenticateRequestWithHeader(r)
	}
	return a.AuthenticateRequestWithCookie(r)
}

// AuthenticateRequestWithHeader validates the JWT from the Authorization header and returns the user ID.
//
// Deprecated: use `Passage.Auth.ValidateJWT` instead.
func (a *App) AuthenticateRequestWithHeader(r *http.Request) (string, error) {
	authHeaderFields := strings.Fields(r.Header.Get("Authorization"))
	if len(authHeaderFields) != 2 || authHeaderFields[0] != "Bearer" {
		return "", Error{Message: "missing authentication token: expected \"Bearer\" header"}
	}

	userID, valid := a.ValidateAuthToken(authHeaderFields[1])
	if !valid {
		return "", Error{Message: "invalid authentication token"}
	}

	return userID, nil
}

// AuthenticateRequestWithCookie validates the JWT from the request cookie and returns the user ID.
//
// Deprecated: use `Passage.Auth.ValidateJWT` instead.
func (a *App) AuthenticateRequestWithCookie(r *http.Request) (string, error) {
	authTokenCookie, err := r.Cookie("psg_auth_token")
	if err != nil {
		return "", Error{Message: "missing authentication token: expected \"psg_auth_token\" cookie"}
	}

	userID, valid := a.ValidateAuthToken(authTokenCookie.Value)
	if !valid {
		return "", Error{Message: "invalid authentication token"}
	}

	return userID, nil
}

// ValidateAuthToken validates the JWT and returns the user ID.
//
// Deprecated: use `Passage.Auth.ValidateJWT` instead.
func (a *App) ValidateAuthToken(authToken string) (string, bool) {
	if authToken == "" {
		return "", false
	}

	parsedToken, err := jwt.Parse(authToken, a.Auth.getPublicKey)
	if err != nil {
		return "", false
	}

	// Extract claims from JWT:
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", false
	}

	if !claims.VerifyAudience(a.ID, true) {
		return "", false
	}

	return userID, true
}

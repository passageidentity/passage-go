package passage

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// AuthenticateRequest determines whether or not to authenticate via header or cookie authentication
// returns the userID (string) on success, error on failure
func (a *App) AuthenticateRequest(r *http.Request) (string, error) {
	if a.Config.HeaderAuth {
		return a.AuthenticateRequestWithHeader(r)
	}
	return a.AuthenticateRequestWithCookie(r)
}

// AuthenticateRequestWithCookie fetches the bearer token from the authorization header and uses it to authenticate
// returns the userID (string) on success, error on failure
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

// getPublicKey returns the associated public key for a JWT
func (a *App) getPublicKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, Error{Message: "expecting JWT header to have string kid"}
	}

	key, ok := a.JWKS.LookupKeyID(keyID)
	// if key doesn't exist, re-fetch one more time to see if this jwk was just added
	if !ok {
		if err := a.refreshJWKSCache(); err != nil {
			return nil, err
		}

		key, ok = a.JWKS.LookupKeyID(keyID)
		if !ok {
			return nil, Error{Message: fmt.Sprintf("unable to find key %q", keyID)}
		}
	}

	var pubKey interface{}
	err := key.Raw(&pubKey)

	return pubKey, err
}

func (a *App) refreshJWKSCache() error {
	var err error
	if a.JWKS, err = a.jwksCache.Refresh(context.Background(), fmt.Sprintf(jwksUrl, a.ID)); err != nil {
		return Error{Message: "failed to fetch jwks"}
	}

	return nil
}

// AuthenticateRequestWithCookie fetches a cookie from the request and uses it to authenticate
// returns the userID (string) on success, error on failure
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

// Deprecated: Use ValidateJwt() instead.
// ValidateAuthToken determines whether a JWT is valid or not
// returns userID (string) on success, error on failure
func (a *App) ValidateAuthToken(authToken string) (string, bool) {
	parsedToken, err := jwt.Parse(authToken, a.getPublicKey)
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

	return userID, true
}

// ValidateJwt determines whether a JWT is valid or not
// returns userID (string) on success, error on failure
func (a *App) ValidateJwt(authToken string) (string, bool) {
	parsedToken, err := jwt.Parse(authToken, a.getPublicKey)
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

	return userID, true
}

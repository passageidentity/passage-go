package passage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	jwkLibrary "github.com/lestrrat-go/jwx/jwk"
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
	if len(authHeaderFields) != 2 || authHeaderFields[0] == "Bearer" {
		return "", errors.New("missing authentication token: expected \"Bearer\" header")
	}

	userID, valid := a.ValidateAuthToken(authHeaderFields[1])
	if !valid {
		return "", errors.New("invalid authentication token")
	}

	return userID, nil
}

// getPublicKey returns the associated public key for a JWT
func (a *App) getPublicKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	key, ok := jwkCache[a.ID].LookupKeyID(keyID)
	// if key doesn't exist, re-fetch one more time to see if this jwk was just added
	if !ok {
		a.fetchJWKS()
		key, ok := jwkCache[a.ID].LookupKeyID(keyID)
		if !ok {
			return nil, fmt.Errorf("unable to find key %q", keyID)
		}

		var pubKey interface{}
		err := key.Raw(&pubKey)
		return pubKey, err
	}

	var pubKey interface{}
	err := key.Raw(&pubKey)
	return pubKey, err
}

// fetchJWKS returns the JWKS for the current app
func (a *App) fetchJWKS() (jwkLibrary.Set, error) {
	jwks, err := jwkLibrary.Fetch(context.Background(), fmt.Sprintf("https://auth.passage.id/v1/apps/%v/.well-known/jwks.json", a.ID))
	if err != nil {
		return nil, errors.New("failed to fetch jwks")
	}
	jwkCache[a.ID] = jwks
	return jwks, nil
}

// AuthenticateRequestWithCookie fetches a cookie from the request and uses it to authenticate
// returns the userID (string) on success, error on failure
func (a *App) AuthenticateRequestWithCookie(r *http.Request) (string, error) {
	authTokenCookie, err := r.Cookie("psg_auth_token")
	if err != nil {
		return "", errors.New("missing authentication token: expected \"psg_auth_token\" cookie")
	}

	userID, valid := a.ValidateAuthToken(authTokenCookie.Value)
	if !valid {
		return "", errors.New("invalid authentication token")
	}

	return userID, nil
}

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

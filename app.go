package passage

import (
	"crypto/rsa"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type App struct {
	Handle    string
	PublicKey *rsa.PublicKey
}

func New() *App {
	return &App{}
}

func (a *App) AuthenticateRequest(r *http.Request) (*User, error) {
	// Check if the app's public key is set. If not, attempt to set it.
	if a.PublicKey == nil {
		pk, err := getPublicKeyFromEnv()
		if err != nil {
			return nil, err
		}
		a.PublicKey = pk
	}

	// Extract authentication token from the request.
	authToken, err := getAuthTokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	// Verify that the authentication token is valid
	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid authentication token")
		}
		return a.PublicKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from JWT
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid authentication token")
	}
	userHandle, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid authentication token")
	}

	// Build a User struct from JWT claims
	user := User{
		Handle: userHandle,
	}

	return &user, nil
}

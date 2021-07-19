package passage

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type App struct {
	apiKey    string
	handle    string
	publicKey *rsa.PublicKey
}

var publicKeyCache map[string]*rsa.PublicKey = make(map[string]*rsa.PublicKey)

func New(appHandle string, params ...string) (*App, error) {
	var apiKey string
	if len(params) > 0 {
		apiKey = params[0]
	}

	var publicKey *rsa.PublicKey
	if cachedPublicKey, ok := publicKeyCache[appHandle]; ok {
		publicKey = cachedPublicKey
	} else {
		var err error
		publicKey, err = fetchPublicKey(appHandle)
		if err != nil {
			return nil, err
		}
		publicKeyCache[appHandle] = publicKey
	}

	return &App{
		apiKey:    apiKey,
		handle:    appHandle,
		publicKey: publicKey,
	}, nil
}

func fetchPublicKey(appHandle string) (*rsa.PublicKey, error) {
	resp, err := http.Get("https://api.passage.id/v1/app/" + appHandle)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body httpResponseError
		err := json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, errors.New("malformatted JSON response")
		}
		return nil, errors.New(body.Message)
	}

	var body struct {
		PublicKey string `json:"public_key"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, errors.New("malformatted JSON response")
	}

	publicKey, err := decodeRSAPublicKey(body.PublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func (a *App) AuthenticateRequest(r *http.Request) (string, error) {
	// Check if the app's public key is set. If not, attempt to set it.
	if a.publicKey == nil {
		return "", errors.New("public key never initialized in app struct")
	}

	// Extract authentication token from the request.
	authToken, err := getAuthTokenFromRequest(r)
	if err != nil {
		return "", err
	}

	// Verify that the authentication token is valid
	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid authentication token")
		}
		return a.publicKey, nil
	})
	if err != nil {
		return "", errors.New("invalid authentication token")
	}

	// Extract claims from JWT
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid authentication token")
	}
	userHandle, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid authentication token")
	}

	return userHandle, nil
}

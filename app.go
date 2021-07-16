package passage

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type App struct {
	handle    string
	publicKey *rsa.PublicKey
	apiKey    string
}
type respErr struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var publicKeyCache map[string]*rsa.PublicKey = make(map[string]*rsa.PublicKey)

func New(appHandle string, key ...string) (*App, error) {
	var apiKeyArg string
	if len(key) > 0 {
		apiKeyArg = key[0]
	}
	var pubKey *rsa.PublicKey
	if cachedPublicKey, ok := publicKeyCache[appHandle]; ok {
		pubKey = cachedPublicKey
	} else {
		resp, err := getAppInfo(appHandle)
		if err != nil {
			return nil, err
		}
		pubKey, err = getRSAPublicKey(resp.PublicKey)
		if err != nil {
			return nil, err
		}
		publicKeyCache[appHandle] = pubKey
	}
	return &App{
		handle:    appHandle,
		publicKey: pubKey,
		apiKey:    apiKeyArg,
	}, nil
}

type appInfoResp struct {
	PublicKey string `json:"public_key"`
}

func getAppInfo(appHandle string) (*appInfoResp, error) {
	resp, err := http.Get("https://api.passage.id/v1/app/" + appHandle)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var respErr respErr
		json.NewDecoder(resp.Body).Decode(&respErr)
		return nil, errors.New(respErr.Message)
	}
	var retBody appInfoResp
	json.NewDecoder(resp.Body).Decode(&retBody)
	return &retBody, nil
}
func (a *App) AuthenticateRequest(r *http.Request) (string, error) {
	// Check if the app's public key is set. If not, attempt to set it.
	if a.publicKey == nil {
		return "", errors.New("public key never initializd in app Struct")
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

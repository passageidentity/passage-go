package passage

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type App struct {
	handle    string
	publicKey *rsa.PublicKey
}

var apiKey string
var publicKeyCache map[string]*rsa.PublicKey = make(map[string]*rsa.PublicKey)

func New(appHandle string, key ...string) (*App, error) {
	if len(key) > 0 {
		apiKey = key[0]
	}
	var pubKey *rsa.PublicKey
	if _, ok := publicKeyCache[appHandle]; ok {
		pubKey = publicKeyCache[appHandle]
	} else {
		var pubKey *rsa.PublicKey
		//MAKE SOME REQUEST TO PASSAGE
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
	}, nil
}

type appInfoResp struct {
	PublicKey string `json:"public_key"`
}

func getAppInfo(appHandle string) (*appInfoResp, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.passage.id/v1/app/"+appHandle, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var retBody appInfoResp

	jsonErr := json.Unmarshal(body, &retBody)
	if jsonErr != nil {
		return nil, err
	}
	return &retBody, nil
}
func (a *App) AuthenticateRequest(r *http.Request) (*User, error) {
	// Check if the app's public key is set. If not, attempt to set it.
	if a.publicKey == nil {
		return nil, errors.New("public key never initializd in app Struct")
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
		return a.publicKey, nil
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

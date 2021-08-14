package passage

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/resty.v1"
)

type Config struct {
	APIKey     string
	CookieAuth bool
}

type App struct {
	ID        string
	PublicKey *rsa.PublicKey
	Config    *Config
}

func New(appID string, config *Config) (*App, error) {
	if config == nil {
		config = &Config{}
	}

	app := App{
		ID:     appID,
		Config: config,
	}

	// Lookup the public key for this App:
	var err error
	app.PublicKey, err = getPublicKey(appID)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

var publicKeyCache map[string]*rsa.PublicKey = make(map[string]*rsa.PublicKey)

func getPublicKey(appID string) (*rsa.PublicKey, error) {
	// First, check if the App's public key is cached locally:
	if cachedPublicKey, ok := publicKeyCache[appID]; ok {
		return cachedPublicKey, nil
	}

	// If the public key isn't cached locally, we'll need to use the Passage API to lookup the public key:
	var responseData struct {
		PublicKey string `json:"public_key"`
	}
	response, err := resty.New().R().
		SetResult(&responseData).
		Get("https://api.passage.id/v1/app/" + appID)
	if err != nil {
		return nil, errors.New("network error: could not lookup Passage App's public key")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("Passage App with ID \"%v\" does not exist", appID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get lookup Passage App's public key")
	}

	// Parse the returned public key string to an rsa.PublicKey:
	publicKeyBytes, err := base64.RawURLEncoding.DecodeString(responseData.PublicKey)
	if err != nil {
		return nil, errors.New("could not parse Passage App's public key: expected valid base-64")
	}
	pemBlock, _ := pem.Decode(publicKeyBytes)
	if pemBlock == nil {
		return nil, errors.New("could not parse Passage App's public key: missing PEM block")
	}
	publicKey, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, errors.New("could not parse Passage App's public key: invalid PKCS #1 public key")
	}

	return publicKey, nil
}

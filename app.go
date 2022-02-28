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
	HeaderAuth bool
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
		App struct {
			PublicKey string `json:"rsa_public_key"`
		} `json:"app"`
	}
	response, err := resty.New().R().
		SetResult(&responseData).
		Get("https://api.passage.id/v1/apps/" + appID)
	if err != nil {
		return nil, errors.New("network error: could not lookup Passage App's public key")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage App with ID \"%v\" does not exist", appID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get lookup Passage App's public key")
	}

	// Parse the returned public key string to an rsa.PublicKey:
	publicKeyBytes, err := base64.RawURLEncoding.DecodeString(responseData.App.PublicKey)
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

type ChannelType string

const (
	EmailChannel ChannelType = "email"
	PhoneChannel ChannelType = "phone"
)

type CreateMagicLinkBody struct {
	UserID        string      `json:"user_id"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	Channel       ChannelType `json:"channel"`
	Send          bool        `json:"send"`
	MagicLinkPath string      `json:"magic_link_path"`
	RedirectURL   string      `json:"redirect_url"`
}

type MagicLink struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	Activated   bool   `json:"activated"`
	UserID      string `json:"user_id"`
	AppID       string `json:"app_id"`
	Identifier  string `json:"identifier"`
	Type        string `json:"type"`
	RedirectURL string `json:"redirect_url"`
	URL         string `json:"url"`
}

// CreateMagicLink receives a CreateMagicLinkBody struct, creating a magic link with provided values
// returns MagicLink on success, error on failure
func (a *App) CreateMagicLink(createMagicLinkBody CreateMagicLinkBody) (*MagicLink, error) {

	type respMagicLink struct {
		MagicLink MagicLink `json:"magic_link"`
	}
	var magicLinkResp respMagicLink

	response, err := resty.New().R().
		SetResult(&magicLinkResp).
		SetBody(&createMagicLinkBody).
		SetAuthToken(a.Config.APIKey).
		Post(fmt.Sprintf("https://api.passage.id/v1/apps/%v/magic-links/", a.ID))
	if err != nil {
		return nil, errors.New("network error: could not create Passage Magic Link")
	}
	if response.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("failed to create Passage Magic Link. Http Status: %v. Response: %v", response.StatusCode(), response.String())
	}

	return &magicLinkResp.MagicLink, nil
}

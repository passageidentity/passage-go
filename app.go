package passage

import (
	"errors"
	"fmt"
	"net/http"

	jwkLibrary "github.com/lestrrat-go/jwx/jwk"
	"gopkg.in/resty.v1"
)

type Config struct {
	APIKey     string
	HeaderAuth bool
}

type App struct {
	ID     string
	JWKS   jwkLibrary.Set
	Config *Config
}

func New(appID string, config *Config) (*App, error) {
	if config == nil {
		config = &Config{}
	}

	app := App{
		ID:     appID,
		Config: config,
	}

	var err error
	app.JWKS, err = app.fetchJWKS()
	if err != nil {
		return nil, err
	}

	return &app, nil
}

var jwkCache map[string]jwkLibrary.Set = make(map[string]jwkLibrary.Set)

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
	TTL           int         `json:"ttl"`
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
	TTL         int    `json:"ttl"`
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

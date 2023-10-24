package passage

import (
	"context"

	jwkLibrary "github.com/lestrrat-go/jwx/jwk"
)

type Config struct {
	APIKey     string
	HeaderAuth bool
}

type App struct {
	ID     string
	JWKS   jwkLibrary.Set
	Config *Config
	client *ClientWithResponses
}

func New(appID string, config *Config) (*App, error) {
	if config == nil {
		config = &Config{}
	}

	client, err := NewClientWithResponses(
		"https://api.passage.id/v1/",
		withPassageVersion,
		withAPIKey(config.APIKey),
	)
	if err != nil {
		return nil, err
	}

	app := App{
		ID:     appID,
		Config: config,
		client: client,
	}

	app.JWKS, err = app.fetchJWKS()
	if err != nil {
		return nil, err
	}

	return &app, nil
}

var jwkCache map[string]jwkLibrary.Set = make(map[string]jwkLibrary.Set)

// GetApp gets information about an app
// returns App on success, error on failure
func (a *App) GetApp() (*AppInfo, error) {
	res, err := a.client.GetAppWithResponse(context.Background(), a.ID)
	if err != nil {
		return nil, Error{Message: "network error: failed to get Passage App Info"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.App, nil
	}

	var errorText string
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    "failed to get Passage App Info",
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// CreateMagicLink receives a CreateMagicLinkBody struct, creating a magic link with provided values
// returns MagicLink on success, error on failure
func (a *App) CreateMagicLink(createMagicLinkBody CreateMagicLinkBody) (*MagicLink, error) {
	res, err := a.client.CreateMagicLinkWithResponse(context.Background(), a.ID, createMagicLinkBody)
	if err != nil {
		return nil, Error{Message: "network error: failed to create Passage Magic Link"}
	}

	if res.JSON201 != nil {
		return &res.JSON201.MagicLink, nil
	}

	var errorText string
	switch {
	case res.JSON400 != nil:
		errorText = res.JSON400.Error
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    "failed to create Passage Magic Link",
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

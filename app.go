package passage

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

const jwksUrl = "https://auth.passage.id/v1/apps/%v/.well-known/jwks.json"

type Config struct {
	APIKey     string
	HeaderAuth bool
}

// Deprecate: will be replace with a different object in v2
type App struct {
	ID        string
	JWKS      jwk.Set
	Config    *Config
	client    *ClientWithResponses
	jwksCache *jwk.Cache
}

// Deprecate: Will be replaced with a different signature in v2
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

	app.jwksCache = jwk.NewCache(context.Background())
	if err := app.jwksCache.Register(fmt.Sprintf(jwksUrl, appID)); err != nil {
		return nil, err
	}

	if err := app.refreshJWKSCache(); err != nil {
		return nil, err
	}

	return &app, nil
}

// Deprecate: GetApp
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

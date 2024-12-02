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

// Deprecated: will be renamed to `Passage` in v2
type App struct {
	ID string
	// Deprecated
	Config       *Config
	User         *appUser
	client       *ClientWithResponses
	jwksCacheSet jwk.Set
}

// Deprecated: Will be replaced with a different signature in v2
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

	url := fmt.Sprintf(jwksUrl, appID)
	cache := jwk.NewCache(context.Background())
	if err := cache.Register(url); err != nil {
		return nil, err
	}

	if _, err = cache.Refresh(context.Background(), url); err != nil {
		return nil, Error{Message: "failed to fetch jwks"}
	}

	app.jwksCacheSet = jwk.NewCachedSet(cache, url)

	app.User = newAppUser(app)

	return &app, nil
}

// GetApp gets information about an app
// returns App on success, error on failure
//
// Deprecated: GetApp - this method will not be replaced
func (a *App) GetApp() (*AppInfo, error) {
	res, err := a.client.GetAppWithResponse(context.Background(), a.ID)
	if err != nil {
		return nil, Error{Message: "network error: failed to get Passage App Info"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.App, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	}

	return nil, Error{
		Message:    "failed to get Passage App Info",
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
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
	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorText = res.JSON400.Error
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	}

	return nil, Error{
		Message:    "failed to create Passage Magic Link",
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

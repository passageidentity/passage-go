package passage

import (
	"context"
	"errors"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

const jwksUrl = "https://auth.passage.id/v1/apps/%v/.well-known/jwks.json"

type Passage = App

// Config holds the configuration for the Passage SDK.
//
// Deprecated: will be removed in v2.
type Config struct {
	APIKey     string
	HeaderAuth bool
}

// App is the main struct for the Passage SDK.
//
// Deprecated: will be renamed to `Passage` in v2.
type App struct {
	// Deprecated: will be removed in v2.
	ID string
	// Deprecated: will be removed in v2.
	Config *Config
	Auth   *auth
	User   *user
	client *ClientWithResponses
}

// New creates a new Passage instance.
//
// Deprecated: Will be replaced with a different signature in v2 -- `New(appID string, apiKey string) (*Passage, error)`.
func New(appID string, config *Config) (*Passage, error) {
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

	url := fmt.Sprintf(jwksUrl, appID)
	cache := jwk.NewCache(context.Background())
	if err := cache.Register(url); err != nil {
		return nil, err
	}

	if _, err = cache.Refresh(context.Background(), url); err != nil {
		return nil, Error{Message: "failed to fetch jwks"}
	}

	auth, err := newAuth(appID, client)
	if err != nil {
		return nil, err
	}
	user := newUser(appID, client)

	return &App{
		ID:     appID,
		Config: config,
		Auth:   auth,
		User:   user,
		client: client,
	}, nil
}

// GetApp fetches the Passage app info.
//
// Deprecated: will be removed in v2.
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

// CreateMagicLink creates a Magic Link for your app.
//
// Deprecated: use `Passage.Auth.CreateMagicLink` instead.
func (a *App) CreateMagicLink(createMagicLinkBody CreateMagicLinkBody) (*MagicLink, error) {
	magicLink, err := a.Auth.CreateMagicLink(createMagicLinkBody)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return magicLink, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return magicLink, err
}

package passage

import (
	"context"
	"fmt"
	"net/http"

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
		withPassageVersion(),
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
		client: client,
		User:   user,
		Auth:   auth,
	}, nil
}

func withPassageVersion() ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Passage-Version", fmt.Sprintf("passage-go %s", version))
		return nil
	})
}

func withAPIKey(apiKey string) ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		if apiKey != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		}
		return nil
	})
}

package passage

import (
	"context"
	"fmt"
	"net/http"
)

// Passage is the main struct for the Passage SDK.
type Passage struct {
	Auth *auth
	User *user
}

// New creates a new Passage instance.
func New(appID string, apiKey string) (*Passage, error) {
	client, err := NewClientWithResponses(
		"https://api.passage.id/v1/",
		withPassageVersion(),
		withAPIKey(apiKey),
	)
	if err != nil {
		return nil, err
	}

	auth, err := newAuth(appID, client)
	if err != nil {
		return nil, err
	}

	user := newUser(appID, client)

	return &Passage{
		User: user,
		Auth: auth,
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

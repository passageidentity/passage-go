package passage

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
)

var withPassageVersion ClientOption = WithRequestEditorFn(
	func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Passage-Version", fmt.Sprintf("passage-go %s", version))
		return nil
	})

func withAPIKey(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		if apiKey != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		}
		return nil
	})
}

package passage

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type auth struct {
	appID        string
	jwksCacheSet jwk.Set
	client       *ClientWithResponses
}

func newAuth(appID string, client *ClientWithResponses) (*auth, error) {
	ctx := context.Background()

	url := fmt.Sprintf(jwksUrl, appID)
	cache := jwk.NewCache(ctx)
	if err := cache.Register(url); err != nil {
		return nil, err
	}

	if _, err := cache.Refresh(ctx, url); err != nil {
		return nil, fmt.Errorf("Failed to fetch JWKS: %w", err)
	}

	auth := auth{
		jwksCacheSet: jwk.NewCachedSet(cache, url),
		client:       client,
	}

	return &auth, nil
}

// CreateMagicLink creates a Magic Link for your app.
func (a *auth) CreateMagicLink(createMagicLinkBody CreateMagicLinkBody) (*MagicLink, error) {
	res, err := a.client.CreateMagicLinkWithResponse(context.Background(), a.appID, createMagicLinkBody)
	if err != nil {
		return nil, Error{Message: "network error: failed to create Passage Magic Link"}
	}

	if res.JSON201 != nil {
		return &res.JSON201.MagicLink, nil
	}

	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    "failed to create Passage Magic Link",
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// ValidateJWT validates the JWT and returns the user ID.
func (a *auth) ValidateJWT(authToken string) (string, error) {
	parsedToken, err := jwt.Parse(authToken, a.getPublicKey)
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Failed to extract claims from JWT")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("Failed to find sub claim in JWT")
	}

	return userID, nil
}

func (a *auth) getPublicKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("Failed to find kid in JWT header")
	}

	key, ok := a.jwksCacheSet.LookupKeyID(keyID)
	if !ok {
		return nil, fmt.Errorf("Failed to find key %q in JWKS", keyID)
	}

	var pubKey interface{}
	err := key.Raw(&pubKey)

	return pubKey, err
}

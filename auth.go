package passage

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	gojwt "github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type auth struct {
	appID        string
	client       *ClientWithResponses
	jwksCacheSet jwk.Set
}

func newAuth(appID string, client *ClientWithResponses) (*auth, error) {
	ctx := context.Background()

	url := fmt.Sprintf(jwksUrl, appID)
	cache := jwk.NewCache(ctx)
	if err := cache.Register(url); err != nil {
		return nil, err
	}

	if _, err := cache.Refresh(ctx, url); err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	auth := auth{
		appID:        appID,
		client:       client,
		jwksCacheSet: jwk.NewCachedSet(cache, url),
	}

	return &auth, nil
}

// CreateMagicLink creates a Magic Link for your app.
func (a *auth) CreateMagicLink(args CreateMagicLinkBody) (*MagicLink, error) {
	res, err := a.client.CreateMagicLinkWithResponse(context.Background(), a.appID, args)
	if err != nil {
		return nil, err
	}

	if res.JSON201 != nil {
		return &res.JSON201.MagicLink, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// ValidateJWT validates the JWT and returns the user ID.
func (a *auth) ValidateJWT(jwt string) (string, error) {
	if jwt == "" {
		return "", errors.New("jwt is required.")
	}

	parsedToken, err := gojwt.Parse(jwt, a.getPublicKey)
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(gojwt.MapClaims)
	if !ok {
		return "", errors.New("failed to extract claims from JWT")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("failed to find sub claim in JWT")
	}

	if !claims.VerifyAudience(a.appID, true) {
		return "", errors.New("failed audience verification for JWT")
	}

	return userID, nil
}

func (a *auth) getPublicKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("failed to find kid in JWT header")
	}

	key, ok := a.jwksCacheSet.LookupKeyID(keyID)
	if !ok {
		return nil, fmt.Errorf("failed to find key %q in JWKS", keyID)
	}

	var pubKey interface{}
	err := key.Raw(&pubKey)

	return pubKey, err
}

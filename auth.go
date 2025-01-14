package passage

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/golang-jwt/jwt"
	gojwt "github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type MagicLinkOptions struct {
	Language      MagicLinkLanguage
	MagicLinkPath string
	RedirectURL   string
	TTL           int
}

type Auth struct {
	appID        string
	client       *ClientWithResponses
	jwksCacheSet jwk.Set
}

func newAuth(appID string, client *ClientWithResponses) (*Auth, error) {
	ctx := context.Background()

	url := fmt.Sprintf("https://auth.passage.id/v1/apps/%v/.well-known/jwks.json", appID)
	cache := jwk.NewCache(ctx)
	if err := cache.Register(url); err != nil {
		return nil, err
	}

	if _, err := cache.Refresh(ctx, url); err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	return &Auth{
		appID:        appID,
		client:       client,
		jwksCacheSet: jwk.NewCachedSet(cache, url),
	}, nil
}

// CreateMagicLink creates a Magic Link for your app using an email address.
func (a *Auth) CreateMagicLinkWithEmail(
	email string,
	magicLinkType MagicLinkType,
	send bool,
	opts *MagicLinkOptions,
) (*MagicLink, error) {
	args := magicLinkArgs{
		Email:       email,
		ChannelType: EmailChannel,
		Type:        magicLinkType,
		Send:        send,
	}

	return a.createMagicLink(args, opts)
}

// CreateMagicLink creates a Magic Link for your app using an E164-formatted phone number.
func (a *Auth) CreateMagicLinkWithPhone(
	phone string,
	magicLinkType MagicLinkType,
	send bool,
	opts *MagicLinkOptions,
) (*MagicLink, error) {
	args := magicLinkArgs{
		Phone:       phone,
		ChannelType: PhoneChannel,
		Type:        magicLinkType,
		Send:        send,
	}

	return a.createMagicLink(args, opts)
}

// CreateMagicLink creates a Magic Link for your app using a Passage user ID.
func (a *Auth) CreateMagicLinkWithUser(
	userID string,
	channel ChannelType,
	magicLinkType MagicLinkType,
	send bool,
	opts *MagicLinkOptions,
) (*MagicLink, error) {
	args := magicLinkArgs{
		UserID:      userID,
		ChannelType: channel,
		Type:        magicLinkType,
		Send:        send,
	}

	return a.createMagicLink(args, opts)
}

// ValidateJWT validates the JWT and returns the user ID.
func (a *Auth) ValidateJWT(jwt string) (string, error) {
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

func (a *Auth) createMagicLink(args magicLinkArgs, opts *MagicLinkOptions) (*MagicLink, error) {
	if opts != nil {
		if err := validateLanguage(opts.Language); err != nil {
			return nil, err
		}

		args.Language = opts.Language
		args.MagicLinkPath = opts.MagicLinkPath
		args.RedirectURL = opts.RedirectURL
		args.TTL = opts.TTL
	}

	res, err := a.client.CreateMagicLinkWithResponse(context.Background(), a.appID, args)
	if err != nil {
		return nil, err
	}

	if res.JSON201 != nil {
		return &res.JSON201.MagicLink, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

func (a *Auth) getPublicKey(token *jwt.Token) (interface{}, error) {
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

func validateLanguage(language MagicLinkLanguage) error {
	if language == "" {
		return nil
	}

	validLanguages := []MagicLinkLanguage{De, En, Es, It, Pl, Pt, Zh}
	if slices.Contains(validLanguages, language) {
		return nil
	}

	return fmt.Errorf("Language must be one of %v", validLanguages)
}

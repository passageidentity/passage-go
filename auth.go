package passage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
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

	rcClient := httprc.NewClient(httprc.WithHTTPClient(http.DefaultClient))

	url := fmt.Sprintf("https://auth.passage.id/v1/apps/%v/.well-known/jwks.json", appID)

	cache, err := jwk.NewCache(ctx, rcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWK cache: %w", err)
	}

	if err := cache.Register(ctx, url); err != nil {
		return nil, fmt.Errorf("failed to register JWKS URL %q in cache: %w", url, err)
	}

	jwksCacheSet, err := cache.Refresh(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch initial JWKS from %q: %w", url, err)
	}

	return &Auth{
		appID:        appID,
		client:       client,
		jwksCacheSet: jwksCacheSet,
	}, nil
}

// CreateMagicLinkWithEmail creates a Magic Link for your app using an email address.
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

// CreateMagicLinkWithPhone creates a Magic Link for your app using an E164-formatted phone number.
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

// CreateMagicLinkWithUser creates a Magic Link for your app using a Passage user ID.
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
func (a *Auth) ValidateJWT(jwtTokenStr string) (string, error) {
	if jwtTokenStr == "" {
		return "", errors.New("jwt is required")
	}

	parsedToken, err := gojwt.Parse(jwtTokenStr, a.getPublicKey)
	if err != nil {
		return "", err // This error could be from parsing, signature validation, or standard claim validation (exp, nbf, iat)
	}

	claims, ok := parsedToken.Claims.(gojwt.MapClaims)
	if !ok {
		return "", errors.New("failed to extract claims from JWT")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("failed to find sub claim in JWT")
	}

	auds, err := claims.GetAudience()
	if err != nil {
		return "", err
	}

	if !slices.Contains(auds, a.appID) {
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

// getPublicKey is the key function for jwt.Parse
// It now correctly uses *gojwt.Token to match the imported package alias
func (a *Auth) getPublicKey(token *gojwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("failed to find kid in JWT header")
	}

	key, ok := a.jwksCacheSet.LookupKeyID(keyID)
	if !ok {
		return nil, fmt.Errorf("failed to find key %q in JWKS", keyID)
	}

	pubKey, err := jwk.PublicRawKeyOf(key)
	if err != nil {
		return nil, fmt.Errorf("failed to extract raw public key: %w", err)
	}
	return pubKey, nil
}

func validateLanguage(language MagicLinkLanguage) error {
	if language == "" {
		return nil
	}

	// Assuming MagicLinkLanguage constants like De, En, Es, etc. are defined elsewhere
	validLanguages := []MagicLinkLanguage{De, En, Es, It, Pl, Pt, Zh}
	if slices.Contains(validLanguages, language) {
		return nil
	}

	return fmt.Errorf("language must be one of %v", validLanguages)
}

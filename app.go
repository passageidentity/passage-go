package passage

import (
	"errors"
	"fmt"
	"net/http"

	jwkLibrary "github.com/lestrrat-go/jwx/jwk"
	"gopkg.in/resty.v1"
)

type Config struct {
	APIKey     string
	HeaderAuth bool
}

type App struct {
	ID     string
	JWKS   jwkLibrary.Set
	Config *Config
}

func New(appID string, config *Config) (*App, error) {
	if config == nil {
		config = &Config{}
	}

	app := App{
		ID:     appID,
		Config: config,
	}

	var err error
	app.JWKS, err = app.fetchJWKS()
	if err != nil {
		return nil, err
	}

	return &app, nil
}

var jwkCache map[string]jwkLibrary.Set = make(map[string]jwkLibrary.Set)

type ChannelType string

const (
	EmailChannel ChannelType = "email"
	PhoneChannel ChannelType = "phone"
)

type CreateMagicLinkBody struct {
	UserID        string      `json:"user_id"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	Channel       ChannelType `json:"channel"`
	Send          bool        `json:"send"`
	MagicLinkPath string      `json:"magic_link_path"`
	RedirectURL   string      `json:"redirect_url"`
	TTL           int         `json:"ttl"`
}

type MagicLink struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	Activated   bool   `json:"activated"`
	UserID      string `json:"user_id"`
	AppID       string `json:"app_id"`
	Identifier  string `json:"identifier"`
	Type        string `json:"type"`
	RedirectURL string `json:"redirect_url"`
	TTL         int    `json:"ttl"`
	URL         string `json:"url"`
}

type AppInfo struct {
	Name                       string              `json:"name"`
	ID                         string              `json:"id"`
	AuthOrigin                 string              `json:"auth_origin"`
	RedirectURL                string              `json:"redirect_url"`
	LoginURL                   string              `json:"login_url"`
	PublicKey                  string              `json:"rsa_public_key"`
	AllowedIdentifier          string              `json:"allowed_identifier"`
	RequiredIdentifier         string              `json:"required_identifier"`
	RequireEmailVerification   bool                `json:"require_email_verification"`
	SessionTimeoutLength       int                 `json:"session_timeout_length"`
	Role                       string              `json:"role,omitempty"`
	UserMetadataSchemaResponse []UserMetadataField `json:"user_metadata_schema"`
	Layouts                    Layouts             `json:"layouts"`
}
type UserMetadataField struct {
	Handle       string                `json:"id"`
	FieldName    string                `json:"field_name"`
	FieldType    UserMetadataFieldType `json:"type"`
	FriendlyName string                `json:"friendly_name"`
	Registration bool                  `json:"registration"`
	Profile      bool                  `json:"profile"`
}

type CreateUserMetadataField struct {
	FriendlyName string                `json:"friendly_name,omitempty"`
	FieldType    UserMetadataFieldType `json:"type,omitempty"`
	Registration bool                  `json:"registration,omitempty"`
	Profile      bool                  `json:"profile,omitempty"`
}

type UpdateUserMetadataField struct {
	FriendlyName string `json:"friendly_name,omitempty"`
	Registration *bool  `json:"registration,omitempty"`
	Profile      *bool  `json:"profile,omitempty"`
}

type UserMetadataFieldType string

const (
	StringMD  UserMetadataFieldType = "string"
	BooleanMD UserMetadataFieldType = "boolean"
	NumberMD  UserMetadataFieldType = "integer"
	DateMD    UserMetadataFieldType = "date"
	PhoneMD   UserMetadataFieldType = "phone"
	EmailMD   UserMetadataFieldType = "email"
)

type Layouts struct {
	Registration []LayoutConfig `json:"registration"`
	Profile      []LayoutConfig `json:"profile"`
}

type LayoutConfig struct {
	ID string `json:"id"`
	X  uint   `json:"x"`
	Y  uint   `json:"y"`
	W  uint   `json:"w"`
	H  uint   `json:"h"`
}

// GetApp gets information about an app
// returns App on success, error on failure
func (a *App) GetApp() (*AppInfo, error) {
	type respAppInfo struct {
		App AppInfo `json:"app"`
	}
	var appResp respAppInfo

	response, err := resty.New().R().
		SetResult(&appResp).
		SetAuthToken(a.Config.APIKey).
		Get(fmt.Sprintf("https://api.passage.id/v1/apps/%v", a.ID))
	if err != nil {
		return nil, errors.New("network error: could not get Passage App Info")
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get Passage App Info. Http Status: %v. Response: %v", response.StatusCode(), response.String())
	}
	fmt.Println(appResp)

	return &appResp.App, nil
}

// CreateMagicLink receives a CreateMagicLinkBody struct, creating a magic link with provided values
// returns MagicLink on success, error on failure
func (a *App) CreateMagicLink(createMagicLinkBody CreateMagicLinkBody) (*MagicLink, error) {

	type respMagicLink struct {
		MagicLink MagicLink `json:"magic_link"`
	}
	var magicLinkResp respMagicLink

	response, err := resty.New().R().
		SetResult(&magicLinkResp).
		SetBody(&createMagicLinkBody).
		SetAuthToken(a.Config.APIKey).
		Post(fmt.Sprintf("https://api.passage.id/v1/apps/%v/magic-links/", a.ID))
	if err != nil {
		return nil, errors.New("network error: could not create Passage Magic Link")
	}
	if response.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("failed to create Passage Magic Link. Http Status: %v. Response: %v", response.StatusCode(), response.String())
	}

	return &magicLinkResp.MagicLink, nil
}

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
	Name                       string              `json:"name"`                       // The name of the App
	ID                         string              `json:"id"`                         // The appID of the App
	AuthOrigin                 string              `json:"auth_origin"`                // The url being used for the App's authentication
	RedirectURL                string              `json:"redirect_url"`               // Where users should be redirected on successful authentication
	LoginURL                   string              `json:"login_url"`                  // Where users should attempt to log in
	PublicKey                  string              `json:"rsa_public_key"`             // The PublicKey associated with the app.
	AllowedIdentifier          string              `json:"allowed_identifier"`         // Which identifier(s) are allowed for this app (email, phone, both)
	RequiredIdentifier         string              `json:"required_identifier"`        // Which identifier(s) are require for this app (email, phone, either, both)
	RequireEmailVerification   bool                `json:"require_email_verification"` // If this app require email verification
	SessionTimeoutLength       int                 `json:"session_timeout_length"`     // How long a JWT will last for the app when a user logs in
	UserMetadataSchemaResponse []UserMetadataField `json:"user_metadata_schema"`       // The schema for user_metadata that will be stored about users
	Layouts                    Layouts             `json:"layouts"`                    // The layouts of user_metadata on the register/profile element
}
type UserMetadataField struct {
	Handle       string                `json:"id"`            // Unique id for the user metadata field
	FieldName    string                `json:"field_name"`    // The name that will be used in user requests to create/update user_metadata
	FieldType    UserMetadataFieldType `json:"type"`          // The type of data stored in this field
	FriendlyName string                `json:"friendly_name"` // The human readable name for this field
	Registration bool                  `json:"registration"`  // Whether or not this field will be accepted on user registration
	Profile      bool                  `json:"profile"`       // Whether or not this field can be update via the passage-profile
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
	Registration []LayoutConfig `json:"registration"` // The UI layout for user_metadata in the passage-register/passage-auth element
	Profile      []LayoutConfig `json:"profile"`      // The UI layout for user_metadata in the passage-profile element
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

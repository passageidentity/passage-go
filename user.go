package passage

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type user struct {
	appID  string
	client *ClientWithResponses
}

func newUser(appID string, client *ClientWithResponses) *user {
	return &user{
		appID:  appID,
		client: client,
	}
}

// Get retrieves a user's object using their user ID.
func (u *user) Get(userID string) (*PassageUser, error) {
	if userID == "" {
		return nil, errors.New("userID is required.")
	}

	res, err := u.client.GetUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.PassageUser, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// GetByIdentifier retrieves a user's object using their user identifier.
func (u *user) GetByIdentifier(identifier string) (*PassageUser, error) {
	if identifier == "" {
		return nil, errors.New("identifier is required.")
	}

	limit := 1
	lowerIdentifier := strings.ToLower(identifier)
	res, err := u.client.ListPaginatedUsersWithResponse(
		context.Background(),
		u.appID,
		&ListPaginatedUsersParams{
			Limit:      &limit,
			Identifier: &lowerIdentifier,
		},
	)

	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users

		if len(users) == 0 {
			return nil, PassageError{
				Message:    "Could not find user with that identifier.",
				ErrorCode:  "user_not_found",
				StatusCode: http.StatusNotFound,
			}
		}

		return u.Get(users[0].ID)
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// Activate activates a user using their user ID.
func (u *user) Activate(userID string) (*PassageUser, error) {
	if userID == "" {
		return nil, errors.New("userID is required.")
	}

	res, err := u.client.ActivateUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.PassageUser, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// Deactivate deactivates a user using their user ID.
func (u *user) Deactivate(userID string) (*PassageUser, error) {
	if userID == "" {
		return nil, errors.New("userID is required.")
	}

	res, err := u.client.DeactivateUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.PassageUser, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// Update updates a user.
func (u *user) Update(userID string, options UpdateUserOptions) (*PassageUser, error) {
	if userID == "" {
		return nil, errors.New("userID is required.")
	}

	res, err := u.client.UpdateUserWithResponse(context.Background(), u.appID, userID, options)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.PassageUser, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// Create creates a user.
func (u *user) Create(args CreateUserArgs) (*PassageUser, error) {
	if args.Email == "" && args.Phone == "" {
		return nil, errors.New("At least one of args.Email or args.Phone is required.")
	}

	res, err := u.client.CreateUserWithResponse(context.Background(), u.appID, args)
	if err != nil {
		return nil, err
	}

	if res.JSON201 != nil {
		return &res.JSON201.PassageUser, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// Delete deletes a user using their user ID.
func (u *user) Delete(userID string) error {
	if userID == "" {
		return errors.New("userID is required.")
	}

	res, err := u.client.DeleteUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	return errorFromResponse(res.Body, res.StatusCode())
}

// ListDevices retrieves a user's webauthn devices using their user ID.
func (u *user) ListDevices(userID string) ([]WebAuthnDevices, error) {
	if userID == "" {
		return nil, errors.New("userID is required.")
	}

	res, err := u.client.ListUserDevicesWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	return nil, errorFromResponse(res.Body, res.StatusCode())
}

// RevokeDevice revokes user's webauthn device using their user ID and the device ID.
func (u *user) RevokeDevice(userID string, deviceID string) error {
	if userID == "" {
		return errors.New("userID is required.")
	}

	if deviceID == "" {
		return errors.New("deviceID is required.")
	}

	res, err := u.client.DeleteUserDevicesWithResponse(context.Background(), u.appID, userID, deviceID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	return errorFromResponse(res.Body, res.StatusCode())
}

// RevokeRefreshTokens revokes all of a user's Refresh Tokens using their User ID.
func (u *user) RevokeRefreshTokens(userID string) error {
	if userID == "" {
		return errors.New("userID is required.")
	}

	res, err := u.client.RevokeUserRefreshTokensWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	return errorFromResponse(res.Body, res.StatusCode())
}

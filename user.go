package passage

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type PassageUser = User
type CreateUserArgs = CreateUserBody
type UpdateUserOptions = UpdateBody

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
	res, err := u.client.GetUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// GetByIdentifier retrieves a user's object using their user identifier.
func (u *user) GetByIdentifier(identifier string) (*PassageUser, error) {
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

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// Activate activates a user using their user ID.
func (u *user) Activate(userID string) (*PassageUser, error) {
	res, err := u.client.ActivateUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// Deactivate deactivates a user using their user ID.
func (u *user) Deactivate(userID string) (*PassageUser, error) {
	res, err := u.client.DeactivateUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// Update updates a user.
func (u *user) Update(userID string, options UpdateUserOptions) (*PassageUser, error) {
	res, err := u.client.UpdateUserWithResponse(context.Background(), u.appID, userID, options)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON400 != nil:
		message = res.JSON400.Error
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// Create creates a user.
func (u *user) Create(args CreateUserArgs) (*PassageUser, error) {
	res, err := u.client.CreateUserWithResponse(context.Background(), u.appID, args)
	if err != nil {
		return nil, err
	}

	if res.JSON201 != nil {
		return &res.JSON201.User, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON400 != nil:
		message = res.JSON400.Error
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// Delete deletes a user using their user ID.
func (u *user) Delete(userID string) error {
	res, err := u.client.DeleteUserWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// ListDevices retrieves a user's webauthn devices using their user ID.
func (u *user) ListDevices(userID string) ([]WebAuthnDevices, error) {
	res, err := u.client.ListUserDevicesWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return nil, err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return nil, PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// RevokeDevice revokes user's webauthn device using their user ID and the device ID.
func (u *user) RevokeDevice(userID string, deviceID string) error {
	res, err := u.client.DeleteUserDevicesWithResponse(context.Background(), u.appID, userID, deviceID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

// RevokeRefreshTokens revokes all of a user's Refresh Tokens using their User ID.
func (u *user) RevokeRefreshTokens(userID string) error {
	res, err := u.client.RevokeUserRefreshTokensWithResponse(context.Background(), u.appID, userID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	var message string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		message = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		message = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		message = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	default:
		var errorBody httpErrorBody
		if err := json.Unmarshal(res.Body, &errorBody); err != nil {
			return err
		}

		message = errorBody.Error
		errorCode = errorBody.Code
	}

	return PassageError{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: res.StatusCode(),
	}
}

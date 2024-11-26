package passage

import (
	"context"
	"net/http"
	"strings"
)

type PassageUser = User
type CreateUserArgs = CreateUserBody
type UpdateUserArgs = UpdateBody

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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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
				ErrorCode:  "user_not_found",
				Message:    "Could not find user with that identifier.",
				StatusCode: http.StatusNotFound,
			}
		}

		return u.Get(users[0].ID)
	}

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
		StatusCode: res.StatusCode(),
	}
}

// Update updates a user.
func (u *user) Update(userID string, options UpdateUserArgs) (*PassageUser, error) {
	res, err := u.client.UpdateUserWithResponse(context.Background(), u.appID, userID, options)
	if err != nil {
		return nil, err
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	var message string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
		message = res.JSON400.Error
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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

	var errorText string
	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorText = res.JSON400.Error
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
		errorCode = string(res.JSON500.Code)
	}

	return nil, Error{
		Message:    errorText,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return nil, PassageError{
		ErrorCode:  errorCode,
		Message:    message,
		StatusCode: res.StatusCode(),
	}
}

// RevokeDevice revokes user's webauthn device using their user ID and the device ID.
func (u *user) RevokeDevice(userID, deviceID string) error {
	res, err := u.client.DeleteUserDevicesWithResponse(context.Background(), u.appID, userID, deviceID)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return nil
	}

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return PassageError{
		ErrorCode:  errorCode,
		Message:    message,
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

	var errorCode string
	var message string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		message = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		message = res.JSON500.Error
	}

	return PassageError{
		ErrorCode:  errorCode,
		Message:    message,
		StatusCode: res.StatusCode(),
	}
}

// type userResponse struct {
// 	JSON200 *UserResponse
// 	JSON400 *N400Error
// 	JSON401 *N401Error
// 	JSON404 *N404Error
// 	JSON500 *N500Error
// }

// func getPassageError[T userResponse](response T) PassageError {
// 	var errorCode string
// 	var message string
// 	switch {
// 	case response.JSON401 != nil:
// 		errorCode = string(response.JSON401.Code)
// 		message = response.JSON401.Error
// 	case response.JSON404 != nil:
// 		errorCode = string(response.JSON404.Code)
// 		message = response.JSON404.Error
// 	case response.JSON500 != nil:
// 		errorCode = string(response.JSON500.Code)
// 		message = response.JSON500.Error
// 	}

// 	return PassageError{
// 		ErrorCode:  errorCode,
// 		Message:    message,
// 		StatusCode: response.StatusCode(),
// 	}
// }

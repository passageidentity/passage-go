package passage

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// GetUser gets a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Get() instead.
func (a *App) GetUser(userID string) (*User, error) {
	message := "failed to get Passage User"
	res, err := a.client.GetUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// GetUserByIdentifier gets a user using their identifier
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.GetByIdentifier() instead.
func (a *App) GetUserByIdentifier(identifier string) (*User, error) {
	var errorText string
	var errorCode string
	message := "failed to get Passage User By Identifier"
	limit := 1
	lowerIdentifier := strings.ToLower(identifier)
	res, err := a.client.ListPaginatedUsersWithResponse(
		context.Background(),
		a.ID,
		&ListPaginatedUsersParams{
			Limit:      &limit,
			Identifier: &lowerIdentifier,
		},
	)

	if err != nil {
		return nil, networkError(fmt.Sprintf("message: %s, err: %+v", message, err))
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users
		if len(users) == 0 {
			message = fmt.Sprintf(IdentifierDoesNotExist, identifier)
			return nil, Error{
				Message:    message,
				StatusCode: http.StatusNotFound,
				StatusText: fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound)),
				ErrorText:  "User not found",
				ErrorCode:  "user_not_found",
			}
		}

		return a.GetUser(users[0].ID)
	}

	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// ActivateUser activates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Activate() instead.
func (a *App) ActivateUser(userID string) (*User, error) {
	message := "failed to activate Passage User"
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// DeactivateUser deactivates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Deactivate() instead.
func (a *App) DeactivateUser(userID string) (*User, error) {
	message := "failed to deactivate Passage User"
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Update() instead.
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {
	message := "failed to update Passage User's attributes"
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.ID, userID, updateBody)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
		errorText = res.JSON400.Error
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
//
// Deprecated: Use Passage.User.Delete() instead.
func (a *App) DeleteUser(userID string) (bool, error) {
	message := "failed to delete Passage User"
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, networkError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// CreateUser receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Create() instead.
func (a *App) CreateUser(createUserBody CreateUserBody) (*User, error) {
	message := "failed to create Passage User"
	res, err := a.client.CreateUserWithResponse(context.Background(), a.ID, createUserBody)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON201 != nil {
		return &res.JSON201.User, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
		errorText = res.JSON400.Error
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// ListUserDevices lists a user's devices
// returns a list of devices on success, error on failure
//
// Deprecated: Use Passage.User.ListDevices() instead.
func (a *App) ListUserDevices(userID string) ([]WebAuthnDevices, error) {
	message := "failed to list devices for a Passage User"
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError(message)
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
//
// Deprecated: Use Passage.User.RevokeDevice() instead.
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {
	message := "failed to delete a device for a Passage User"
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.ID, userID, deviceID)
	if err != nil {
		return false, networkError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		switch res.JSON404.Code {
		case UserNotFound:
			message = fmt.Sprintf(UserIDDoesNotExist, userID)
		case DeviceNotFound:
			message = fmt.Sprintf("Device with ID \"%v\" does not exist", deviceID)
		}
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
//
// Deprecated: Use Passage.User.SignOut() instead.
func (a *App) SignOut(userID string) (bool, error) {
	message := "failed to revoke all refresh tokens for a Passage User"
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, networkError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

func networkError(message string) error {
	return Error{
		Message:    message,
		StatusCode: 500,
		StatusText: "500 Service Error",
		ErrorText:  "internal_service_error",
		ErrorCode:  "Internal Service Error",
	}
}

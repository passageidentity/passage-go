package passage

import (
	"context"
	"fmt"
)

const UserIDDoesNotExist string = "passage User with ID \"%v\" does not exist"

// GetUser gets a user using their userID
// returns user on success, error on failure
func (a *App) GetUser(userID string) (*User, error) {
	res, err := a.client.GetUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, Error{Message: "network error: failed to get Passage User"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	message := "failed to get Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// ActivateUser activates a user using their userID
// returns user on success, error on failure
func (a *App) ActivateUser(userID string) (*User, error) {
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, Error{Message: "network error: failed to activate Passage User"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	message := "failed to activate Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// DeactivateUser deactivates a user using their userID
// returns user on success, error on failure
func (a *App) DeactivateUser(userID string) (*User, error) {
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, Error{Message: "network error: failed to deactivate Passage User"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	message := "failed to deactivate Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.ID, userID, updateBody)
	if err != nil {
		return nil, Error{Message: "network error: failed to update Passage User's attributes"}
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	message := "failed to update Passage User's attributes"
	switch {
	case res.JSON400 != nil:
		errorText = res.JSON400.Error
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
func (a *App) DeleteUser(userID string) (bool, error) {
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, Error{Message: "network error: failed to delete Passage User"}
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	message := "failed to delete Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// CreateUser receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *App) CreateUser(createUserBody CreateUserBody) (*User, error) {
	res, err := a.client.CreateUserWithResponse(context.Background(), a.ID, createUserBody)
	if err != nil {
		return nil, Error{Message: "network error: failed to create Passage User"}
	}

	if res.JSON201 != nil {
		return &res.JSON201.User, nil
	}

	var errorText string
	switch {
	case res.JSON400 != nil:
		errorText = res.JSON400.Error
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    "failed to create Passage User",
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// ListUserDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *App) ListUserDevices(userID string) ([]WebAuthnDevices, error) {
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, Error{Message: "network error: failed to list devices for a Passage User"}
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var errorText string
	message := "failed to list devices for a Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return nil, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.ID, userID, deviceID)
	if err != nil {
		return false, Error{Message: "network error: failed to delete a device for a Passage User"}
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	message := "failed to delete a device for a Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		switch res.JSON404.Code {
		case UserNotFound:
			message = fmt.Sprintf(UserIDDoesNotExist, userID)
		case DeviceNotFound:
			message = fmt.Sprintf("Device with ID \"%v\" does not exist", deviceID)
		}
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
func (a *App) SignOut(userID string) (bool, error) {
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, Error{Message: "network error: failed to revoke all refresh tokens for a Passage User"}
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	message := "failed to revoke all refresh tokens for a Passage User"
	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorText = res.JSON500.Error
	}

	return false, Error{
		Message:    message,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
	}
}

package passage

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type PassageUser = User
type appUser struct {
	client *ClientWithResponses
	appID  string
}

const (
	UserIDDoesNotExist     string = "passage User with ID \"%v\" does not exist"
	IdentifierDoesNotExist string = "passage User with Identifier \"%v\" does not exist"
)

func newAppUser(client *ClientWithResponses, appID string) *appUser {
	appUser := appUser{
		client: client,
		appID:  appID,
	}

	return &appUser
}

// Get gets a user using their userID
// returns user on success, error on failure
func (a *appUser) Get(userID string) (*PassageUser, error) {
	message := "failed to get Passage User"
	res, err := a.client.GetUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// GetByIdentifier gets a user using their identifier
// returns user on success, error on failure
func (a *appUser) GetByIdentifier(identifier string) (*PassageUser, error) {
	var errorCode string
	message := "failed to get Passage User By Identifier"
	limit := 1
	lowerIdentifier := strings.ToLower(identifier)
	res, err := a.client.ListPaginatedUsersWithResponse(
		context.Background(),
		a.appID,
		&ListPaginatedUsersParams{
			Limit:      &limit,
			Identifier: &lowerIdentifier,
		},
	)

	if err != nil {
		return nil, networkPassageError(fmt.Sprintf("message: %s, err: %+v", message, err))
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users
		if len(users) == 0 {
			message = fmt.Sprintf(IdentifierDoesNotExist, identifier)
			return nil, PassageError{
				Message:    message,
				StatusCode: http.StatusNotFound,
				ErrorCode:  "user_not_found",
			}
		}

		return a.Get(users[0].ID)
	}

	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// Activate activates a user using their userID
// returns user on success, error on failure
func (a *appUser) Activate(userID string) (*PassageUser, error) {
	message := "failed to activate Passage User"
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// Deactivate deactivates a user using their userID
// returns user on success, error on failure
func (a *appUser) Deactivate(userID string) (*PassageUser, error) {
	message := "failed to deactivate Passage User"
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// Update receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *appUser) Update(userID string, updateBody UpdateBody) (*PassageUser, error) {
	message := "failed to update Passage User's attributes"
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.appID, userID, updateBody)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// Delete deletes a user by their user string
// returns true on success, false and error on failure (bool, err)
func (a *appUser) Delete(userID string) (bool, error) {
	message := "failed to delete Passage User"
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return false, networkPassageError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return false, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// Create receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *appUser) Create(createUserBody CreateUserBody) (*PassageUser, error) {
	message := "failed to create Passage User"
	res, err := a.client.CreateUserWithResponse(context.Background(), a.appID, createUserBody)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON201 != nil {
		return &res.JSON201.User, nil
	}

	var errorCode string
	switch {
	case res.JSON400 != nil:
		errorCode = string(res.JSON400.Code)
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// ListDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *appUser) ListDevices(userID string) ([]WebAuthnDevices, error) {
	message := "failed to list devices for a Passage User"
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError(message)
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return nil, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// RevokeDevice gets a user using their userID
// returns a true success, error on failure
func (a *appUser) RevokeDevice(userID, deviceID string) (bool, error) {
	message := "failed to delete a device for a Passage User"
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.appID, userID, deviceID)
	if err != nil {
		return false, networkPassageError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		switch res.JSON404.Code {
		case UserNotFound:
			message = fmt.Sprintf(UserIDDoesNotExist, userID)
		case DeviceNotFound:
			message = fmt.Sprintf("Device with ID \"%v\" does not exist", deviceID)
		}
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return false, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

// RevokeRefreshTokens revokes a users refresh tokens
// returns true on success, error on failure
func (a *appUser) RevokeRefreshTokens(userID string) (bool, error) {
	message := "failed to revoke all refresh tokens for a Passage User"
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return false, networkPassageError(message)
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	switch {
	case res.JSON401 != nil:
		errorCode = string(res.JSON401.Code)
	case res.JSON404 != nil:
		errorCode = string(res.JSON404.Code)
		message = fmt.Sprintf(UserIDDoesNotExist, userID)
	case res.JSON500 != nil:
		errorCode = string(res.JSON500.Code)
	}

	return false, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

func networkPassageError(message string) PassageError {
	return PassageError{
		Message:    message,
		StatusCode: 500,
		ErrorCode:  "Internal Service Error",
	}
}

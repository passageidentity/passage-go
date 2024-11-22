package passage

import (
	"context"
	"net/http"
	"strings"
)

type PassageUser = User
type appUser struct {
	client *ClientWithResponses
	appID  string
}

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
	res, err := a.client.GetUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	var message string
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
	var message string
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
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users
		if len(users) == 0 {
			return nil, PassageError{
				Message:    "User not found",
				StatusCode: http.StatusNotFound,
				ErrorCode:  "user_not_found",
			}
		}

		return a.Get(users[0].ID)
	}

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
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.appID, userID, updateBody)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return false, networkPassageError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.CreateUserWithResponse(context.Background(), a.appID, createUserBody)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON201 != nil {
		return &res.JSON201.User, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return nil, networkPassageError()
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.appID, userID, deviceID)
	if err != nil {
		return false, networkPassageError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	var message string
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
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.appID, userID)
	if err != nil {
		return false, networkPassageError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorCode string
	var message string
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
	}

	return false, PassageError{
		Message:    message,
		StatusCode: res.StatusCode(),
		ErrorCode:  errorCode,
	}
}

func networkPassageError() PassageError {
	return PassageError{
		Message:    "Internal Service Error",
		StatusCode: 500,
		ErrorCode:  "internal_service_error",
	}
}

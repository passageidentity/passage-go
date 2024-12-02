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
	res, err := a.client.GetUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
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

// GetUserByIdentifier gets a user using their identifier
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.GetByIdentifier() instead.
func (a *App) GetUserByIdentifier(identifier string) (*User, error) {
	var errorText string
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
		return nil, networkError()
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users
		if len(users) == 0 {
			return nil, Error{
				Message:    "User not found",
				StatusCode: http.StatusNotFound,
				StatusText: fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound)),
				ErrorText:  "User not found",
				ErrorCode:  "user_not_found",
			}
		}

		return a.GetUser(users[0].ID)
	}

	var errorCode string
	switch {
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

// ActivateUser activates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Activate() instead.
func (a *App) ActivateUser(userID string) (*User, error) {
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
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

// DeactivateUser deactivates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Deactivate() instead.
func (a *App) DeactivateUser(userID string) (*User, error) {
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
	}

	var errorText string
	var errorCode string
	switch {
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

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
//
// Deprecated: Use Passage.User.Update() instead.
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.ID, userID, updateBody)
	if err != nil {
		return nil, networkError()
	}

	if res.JSON200 != nil {
		return &res.JSON200.User, nil
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

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
//
// Deprecated: Use Passage.User.Delete() instead.
func (a *App) DeleteUser(userID string) (bool, error) {
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, networkError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
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

	return false, Error{
		Message:    errorText,
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
	res, err := a.client.CreateUserWithResponse(context.Background(), a.ID, createUserBody)
	if err != nil {
		return nil, networkError()
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

// ListUserDevices lists a user's devices
// returns a list of devices on success, error on failure
//
// Deprecated: Use Passage.User.ListDevices() instead.
func (a *App) ListUserDevices(userID string) ([]WebAuthnDevices, error) {
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return nil, networkError()
	}

	if res.JSON200 != nil {
		return res.JSON200.Devices, nil
	}

	var errorText string
	var errorCode string
	switch {
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

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
//
// Deprecated: Use Passage.User.RevokeDevice() instead.
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.ID, userID, deviceID)
	if err != nil {
		return false, networkError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
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

	return false, Error{
		Message:    errorText,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
//
// Deprecated: Use Passage.User.RevokeRefreshTokens() instead.
func (a *App) SignOut(userID string) (bool, error) {
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.ID, userID)
	if err != nil {
		return false, networkError()
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		return true, nil
	}

	var errorText string
	var errorCode string
	switch {
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

	return false, Error{
		Message:    errorText,
		StatusCode: res.StatusCode(),
		StatusText: res.Status(),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}
}

func networkError() error {
	return Error{
		Message:    "Internal Service Error",
		StatusCode: 500,
		StatusText: "500 Service Error",
		ErrorText:  "internal_service_error",
		ErrorCode:  "internal_service_error",
	}
}

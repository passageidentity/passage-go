package passage

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AppUser struct {
	AppID  string
	UserID string
	client *ClientWithResponses
}

const (
	UserIDDoesNotExist     string = "passage User with ID \"%v\" does not exist"
	IdentifierDoesNotExist string = "passage User with Identifier \"%v\" does not exist"
)

func NewAppUser(appID, userID string, config *Config) (*AppUser, error) {
	if config == nil {
		config = &Config{}
	}

	client, err := NewClientWithResponses(
		"https://api.passage.id/v1/",
		withPassageVersion,
		withAPIKey(config.APIKey),
	)
	if err != nil {
		return nil, err
	}

	appUser := AppUser{
		AppID:  appID,
		UserID: userID,
		client: client,
	}

	return &appUser, nil
}

func NewAppUserByIdentifier(appID, identifier string, config *Config) (*AppUser, error) {
	if config == nil {
		config = &Config{}
	}

	client, err := NewClientWithResponses(
		"https://api.passage.id/v1/",
		withPassageVersion,
		withAPIKey(config.APIKey),
	)
	if err != nil {
		return nil, err
	}

	userID, err := getUserIdByIdentifier(appID, identifier, client)

	appUser := AppUser{
		AppID:  appID,
		UserID: *userID,
		client: client,
	}

	return &appUser, nil
}

// Get gets a user using their userID
// returns user on success, error on failure
func (a *AppUser) Get() (*User, error) {
	a.validate()
	res, err := a.client.GetUserWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// Activate activates a user using their userID
// returns user on success, error on failure
func (a *AppUser) Activate() (*User, error) {
	a.validate()
	res, err := a.client.ActivateUserWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// Deactivate deactivates a user using their userID
// returns user on success, error on failure
func (a *AppUser) Deactivate() (*User, error) {
	a.validate()
	res, err := a.client.DeactivateUserWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// Update receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *AppUser) Update(updateBody UpdateBody) (*User, error) {
	a.validate()
	res, err := a.client.UpdateUserWithResponse(context.Background(), a.AppID, a.UserID, updateBody)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// Delete deletes a user by their user string
// returns true on success, false and error on failure (bool, err)
func (a *AppUser) Delete() (bool, error) {
	a.validate()
	res, err := a.client.DeleteUserWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// Create receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *AppUser) Create(createUserBody CreateUserBody) (*User, error) {
	a.validateForCreate()
	res, err := a.client.CreateUserWithResponse(context.Background(), a.AppID, createUserBody)
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

// ListDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *AppUser) ListDevices() ([]WebAuthnDevices, error) {
	a.validate()
	res, err := a.client.ListUserDevicesWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// RevokeDevice gets a user using their userID
// returns a true success, error on failure
func (a *AppUser) RevokeDevice(deviceID string) (bool, error) {
	a.validate()
	res, err := a.client.DeleteUserDevicesWithResponse(context.Background(), a.AppID, a.UserID, deviceID)
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
			message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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
func (a *AppUser) SignOut() (bool, error) {
	a.validate()
	res, err := a.client.RevokeUserRefreshTokensWithResponse(context.Background(), a.AppID, a.UserID)
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
		message = fmt.Sprintf(UserIDDoesNotExist, a.UserID)
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

// getByIdentifier gets a userID using their identifier
// returns userID on success, error on failure
func getUserIdByIdentifier(appID, identifier string, client *ClientWithResponses) (*string, error) {
	var errorText string
	message := "failed to get Passage User By Identifier"
	limit := 1
	lowerIdentifier := strings.ToLower(identifier)
	res, err := client.ListPaginatedUsersWithResponse(
		context.Background(),
		appID,
		&ListPaginatedUsersParams{
			Limit:      &limit,
			Identifier: &lowerIdentifier,
		},
	)

	if err != nil {
		return nil, Error{Message: fmt.Sprintf("network error:failed to get Passage User by Identifier. message: %s, err: %+v", message, err)}
	}

	if res.JSON200 != nil {
		users := res.JSON200.Users
		if len(users) == 0 {
			message = fmt.Sprintf(IdentifierDoesNotExist, identifier)
			return nil, Error{
				Message:    message,
				StatusCode: http.StatusNotFound,
				StatusText: http.StatusText(http.StatusNotFound),
				ErrorText:  "User not found",
			}
		}

		userID := users[0].ID
		return &userID, nil
	}

	switch {
	case res.JSON401 != nil:
		errorText = res.JSON401.Error
	case res.JSON404 != nil:
		errorText = res.JSON404.Error
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

func (a *AppUser) validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.AppID, validation.Required),
		validation.Field(&a.client, validation.Required),
		validation.Field(&a.UserID, validation.Required),
	)
}

func (a *AppUser) validateForCreate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.AppID, validation.Required),
		validation.Field(&a.client, validation.Required),
	)
}

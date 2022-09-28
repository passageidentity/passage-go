package passage

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type UserStatus string

const (
	UserIDDoesNotExist string     = "passage User with ID \"%v\" does not exist"
	StatusActive       UserStatus = "active"
	StatusInactive     UserStatus = "inactive"
	StatusPending      UserStatus = "pending"
)

type User struct {
	ID            string                 `json:"id"`
	Status        UserStatus             `json:"status"`
	Email         string                 `json:"email"`
	Phone         string                 `json:"phone"`
	EmailVerified bool                   `json:"email_verified"`
	PhoneVerified bool                   `json:"phone_verified"`
	CreatedAt     time.Time              `json:"created_at"`
	LastLogin     time.Time              `json:"last_login_at"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
}

type Device struct {
	ID          string    `json:"id"`
	CredID      string    `json:"cred_id"`
	Name        string    `json:"friendly_name"`
	UsageCount  uint      `json:"usage_count"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

// GetUser gets a user using their userID
// returns user on success, error on failure
func (a *App) GetUser(userID string) (*User, error) {
	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		SetError(&errorResponse).
		Get(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return nil, Error{Message: "network error: failed to get Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return nil, Error{
			Message:    "failed to get Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	user := userBody.User

	return &user, nil
}

// ActivateUser activates a user using their userID
// returns user on success, error on failure
func (a *App) ActivateUser(userID string) (*User, error) {
	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		SetError(&errorResponse).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/activate", a.ID, userID))
	if err != nil {
		return nil, Error{Message: "network error: failed to get activate Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return nil, Error{
			Message:    "failed to activate Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	user := userBody.User

	return &user, nil
}

// DeactivateUser deactivates a user using their userID
// returns user on success, error on failure
func (a *App) DeactivateUser(userID string) (*User, error) {
	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		SetBody(&errorResponse).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/deactivate", a.ID, userID))
	if err != nil {
		return nil, Error{Message: "network error: failed to deactivate Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return nil, Error{
			Message:    "failed to deactivate Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	user := userBody.User

	return &user, nil
}

type UpdateBody struct {
	Email        string                 `json:"email,omitempty"`
	Phone        string                 `json:"phone,omitempty"`
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
}

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {

	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		SetBody(updateBody).
		SetError(&errorResponse).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return nil, Error{Message: "network error: failed to update Passage User attributes"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return nil, Error{
			Message:    "failed to update Passage User's attributes",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	user := userBody.User

	return &user, nil
}

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
func (a *App) DeleteUser(userID string) (bool, error) {

	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetError(&errorResponse).
		Delete(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return false, Error{Message: "network error: could not delete Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return false, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return false, Error{
			Message:    "failed to delete Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}

	return true, nil
}

type CreateUserBody struct {
	Email        string                 `json:"email,omitempty"`
	Phone        string                 `json:"phone,omitempty"`
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
}

// CreateUser receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *App) CreateUser(createUserBody CreateUserBody) (*User, error) {

	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetResult(&userBody).
		SetBody(&createUserBody).
		SetError(&errorResponse).
		SetAuthToken(a.Config.APIKey).
		Post(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/", a.ID))
	if err != nil {
		return nil, Error{Message: "network error: failed create Passage User"}
	}
	if response.StatusCode() != http.StatusCreated {
		return nil, Error{
			Message:    "failed to create Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	user := userBody.User

	return &user, nil
}

// ListUserDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *App) ListUserDevices(userID string) ([]Device, error) {
	type respDevices struct {
		Devices []Device `json:"devices"`
	}
	var devicesBody respDevices
	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&devicesBody).
		SetError(&errorResponse).
		Get(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/devices", a.ID, userID))
	if err != nil {
		return nil, Error{Message: "network error: failed to list devices for a Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return nil, Error{
			Message:    "failed to list devices for a Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	devices := devicesBody.Devices

	return devices, nil
}

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {

	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetError(&errorResponse).
		Delete(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/devices/%v", a.ID, userID, deviceID))
	if err != nil {
		return false, Error{Message: "network error: failed to delete a device for a Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return false, Error{
			Message:    fmt.Sprintf("passage User with ID \"%v\" does not exist or Devices with ID \"%v\" does not exist", userID, deviceID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return false, Error{
			Message:    "failed to delete a device for a Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}

	return true, nil
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
func (a *App) SignOut(userID string) (bool, error) {

	var errorResponse HTTPError

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetError(&errorResponse).
		Delete(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/tokens/", a.ID, userID))
	if err != nil {
		return false, Error{Message: "network error: failed to revoke all refresh tokens for a Passage User"}
	}
	if response.StatusCode() == http.StatusNotFound {
		return false, Error{
			Message:    fmt.Sprintf(UserIDDoesNotExist, userID),
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}
	if response.StatusCode() != http.StatusOK {
		return false, Error{
			Message:    "failed to revoke all refresh tokens for a Passage User",
			StatusCode: response.StatusCode(),
			StatusText: http.StatusText(response.StatusCode()),
			ErrorText:  errorResponse.ErrorText,
		}
	}

	return true, nil
}

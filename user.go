package passage

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
	StatusPending  UserStatus = "pending"
)

type User struct {
	ID            string     `json:"id"`
	Status        UserStatus `json:"status"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	EmailVerified bool       `json:"email_verified"`
	CreatedAt     time.Time  `json:"created_at"`
	LastLogin     time.Time  `json:"last_login_at"`
}

type Device struct {
	ID         string    `json:"id"`
	CredID     string    `json:"cred_id"`
	Name       string    `json:"friendly_name"`
	UsageCount uint      `json:"usage_count"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

// GetUser gets a user using their userID
// returns user on success, error on failure
func (a *App) GetUser(userID string) (*User, error) {
	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		Get(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not get Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get Passage User")
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

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/activate", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not get activate Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to activate Passage User")
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

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/deactivate", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not get deactivate Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to deactivate Passage User")
	}
	user := userBody.User

	return &user, nil
}

type UpdateBody struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {

	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&userBody).
		SetBody(updateBody).
		Patch(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not update Passage User attributes")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to patch Passage User's attributes")
	}
	user := userBody.User

	return &user, nil
}

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
func (a *App) DeleteUser(userID string) (bool, error) {
	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		Delete(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v", a.ID, userID))
	if err != nil {
		return false, errors.New("network error: could not delete Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return false, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("failed to delete Passage User")
	}

	return true, nil
}

type CreateUserBody struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

// CreateUser receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *App) CreateUser(createUserBody CreateUserBody) (*User, error) {

	type respUser struct {
		User User `json:"user"`
	}
	var userBody respUser

	response, err := resty.New().R().
		SetResult(&userBody).
		SetBody(createUserBody).
		SetAuthToken(a.Config.APIKey).
		Post(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/", a.ID))
	if err != nil {
		return nil, errors.New("network error: could not create Passage User")
	}
	if response.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("failed to create Passage User. Http Status: %v. Response: %v", response.StatusCode(), response.String())
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

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&devicesBody).
		Get(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/devices", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not get Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to list devices for a Passage User")
	}
	devices := devicesBody.Devices

	return devices, nil
}

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {
	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		Delete(fmt.Sprintf("https://api.passage.id/v1/apps/%v/users/%v/devices/%v", a.ID, userID, deviceID))
	if err != nil {
		return false, errors.New("network error: could not get Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return false, fmt.Errorf("passage User with ID \"%v\" does not exist or Devices with ID \"%v\" does not exist", userID, deviceID)
	}
	if response.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("failed to delete a device for a Passage User")
	}

	return true, nil
}

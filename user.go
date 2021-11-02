package passage

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type User struct {
	ID            string    `json:"id"`
	Active        bool      `json:"active"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	LastLogin     time.Time `json:"last_login_at"`
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

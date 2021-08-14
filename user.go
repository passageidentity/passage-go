package passage

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type UserEvents struct {
	EventType string    `json:"event_type"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type User struct {
	ID            string    `json:"handle"`
	Active        bool      `json:"active"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	StartDate     time.Time `json:"start_date"`
	LastLogin     time.Time `json:"last_login_date"`
}

func (a *App) GetUser(userID string) (*User, error) {
	var user User

	response, err := resty.New().R().
		SetAuthToken(a.Config.APIKey).
		SetResult(&user).
		Get(fmt.Sprintf("https://api.passage.id/v1/app/%v/users/%v", a.ID, userID))
	if err != nil {
		return nil, errors.New("network error: could not get Passage User")
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("Passage User with ID \"%v\" does not exist", userID)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get Passage User")
	}

	return &user, nil
}

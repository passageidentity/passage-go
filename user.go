package passage

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type UserEvents struct {
	EventType string    `json:"type"`
	Id        string    `json:"id"`
	Timestamp time.Time `json:"created_at"`
}

type User struct {
	Active          bool         `json:"active"`
	Email           string       `json:"email"`
	EmailVerified   bool         `json:"email_verified"`
	Id              string       `json:"id"`
	CreatedAt       time.Time    `json:"created_at"`
	LastLoginAt     time.Time    `json:"last_login_at"`
	LoginCount      int          `json:"login_count"`
	RecentEvents    []UserEvents `json:"recent_events"`
	Webauthn        bool         `json:"webauthn"`
	WebauthnDevices []string     `json:"webauthn_devices"`
}

func (a *App) GetUser(userID string) (*User, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, "https://api.passage.id/v1/apps/"+a.id+"/users/"+userID, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+a.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body httpResponseError
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, errors.New("malformatted JSON response")
		}
		return nil, errors.New(body.Message)
	}

	type userBody struct {
		User User `json:"user"`
	}
	var respUser userBody
	err = json.NewDecoder(resp.Body).Decode(&respUser)
	if err != nil {
		return nil, errors.New("malformatted JSON response")
	}

	user := respUser.User

	return &user, nil
}

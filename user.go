package passage

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type UserEvents struct {
	EventType string    `json:"event_type"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type User struct {
	Active          bool         `json:"active"`
	Email           string       `json:"email"`
	EmailVerified   bool         `json:"email_verified"`
	Handle          string       `json:"handle"`
	StartDate       time.Time    `json:"start_date"`
	LastLogin       time.Time    `json:"last_login"`
	RecentEvents    []UserEvents `json:"recent_events"`
	Password        bool         `json:"password"`
	Webauthn        bool         `json:"webauthn"`
	WebauthnDevices []string     `json:"webauthn_devices"`
}

func (a *App) GetUser(userHandle string) (*User, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, "https://api.passage.id/v1/app/"+a.handle+"/users/"+userHandle, nil)
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

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, errors.New("malformatted JSON response")
	}

	return &user, nil
}

package passage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type UserEvents struct {
	EventType string    `json:"event_type"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type User struct {
	StartDate       time.Time    `json:"start_date"`
	Active          bool         `json:"active"`
	EmailVerified   bool         `json:"email_verified"`
	Email           string       `json:"email"`
	Handle          string       `json:"handle"`
	LastLogin       time.Time    `json:"last_login"`
	RecentEvents    []UserEvents `json:"recent_events"`
	Password        bool         `json:"password"`
	Webauthn        bool         `json:"webauthn"`
	WebauthnDevices []string     `json:"webauthn_devices"`
}

func (a *App) GetUserInfo(userHandle string) (*User, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.passage.id/v1/app/"+a.handle+"/users/"+userHandle, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+a.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var retBody User

	jsonErr := json.Unmarshal(body, &retBody)
	if jsonErr != nil {
		return nil, err
	}
	return &retBody, nil
}

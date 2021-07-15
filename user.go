package passage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Handle string
}

func (a *App) GetUserInfo(userHandle string) (interface{}, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.passage.id/v1/app/"+a.handle+"/users/"+userHandle, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var retBody interface{}

	jsonErr := json.Unmarshal(body, &retBody)
	if jsonErr != nil {
		return nil, err
	}
	return retBody, nil
}

package passage

import (
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	//will fail unless API_KEY is replaced with a valid key
	psg, err := New("UKbRUx", "API_KEY")
	if err != nil {
		t.Error(err)
	}
	_, errReq := psg.GetUserInfo("jAOBfYtZNoxVdFGjUwQB")
	if errReq != nil {
		t.Error(errReq)
	}
}

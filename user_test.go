package passage

import (
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	psg, err := New("UKbRUx", "xcNm7N57In.8nigt09Hrx4lf3pqNsxnbAG34h312uWCxk9S7X1zTB608LiNRexgeUXAUeAJWhFB")
	if err != nil {
		t.Error(err)
	}
	_, errReq := psg.GetUserInfo("jAOBfYtZNoxVdFGjUwQB")
	if errReq != nil {
		t.Error(errReq)
	}
}

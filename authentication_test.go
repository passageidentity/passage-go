package passage_test

import (
	"net/http"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationWithCookie(t *testing.T) {
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.Nil(t, err)

	psg, err := passage.New("TrWSUbDDTPCKTQDtLA9MO8Ee", &passage.Config{
		HeaderAuth: false,
	})
	require.Nil(t, err)

	t.Run("valid auth token", func(t *testing.T) {
		req.AddCookie(&http.Cookie{
			Name:  "psg_auth_token",
			Value: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjlGVGdhaWpXV2hBUFhyTmJNQmMxc1lxWCIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwOi8vbG9jYWxob3N0OjUwMDAiLCJleHAiOjE2ODg0OTA5NzcsImlhdCI6MTY1MjIwMjk3NywiaXNzIjoiVHJXU1ViRERUUENLVFFEdExBOU1POEVlIiwibmJmIjoxNjUyMjAyOTc3LCJzdWIiOiJiRVhJWktZeUFwZ3o1b1dZYzVXTTl2ZkYifQ.RFDUgv4ewmeEnapatQONuNJuofuTKDC7r7gZuvPGWpoX_EJWCgjVuysVt4L8ghUO_ZUuaujEn7loSZAtVVG7NKmivN2hSfCtCoK6JW-y8fn3izlaERl5fldkNdN8rxISlgqtANuPV0xfxtbIoqagV9wCAt2DY53HXDYM13ZRHIDrXgRO3-kiPhp_mO_tUnvHBRZ59DDFd-nqk99ssepT0-uEl-KVcHIQKbt5SfGgM9sR-b30mp6g-PkDDgdpMmS-ZLCNAZTkDclHTlCEdxCTdHS46z6yz6QAVgzhU0Z48q6olBzBEMGn-OC0fbwkYjG-j6xyWhpri0BamAG5I-i5Nw",
		})

		userID, err := psg.AuthenticateRequest(req)
		require.Nil(t, err)
		assert.NotNil(t, userID)
	})
}

func TestAuthenticationWithCookiePassage(t *testing.T) {
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.Nil(t, err)

	psg, err := passage.New("passage", &passage.Config{
		HeaderAuth: false,
	})
	require.Nil(t, err)

	t.Run("fail with missing auth token", func(t *testing.T) {
		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})

	t.Run("fail with invalid auth token", func(t *testing.T) {
		req.AddCookie(&http.Cookie{
			Name:  "psg_auth_token",
			Value: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA",
		})

		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})
}

func TestAuthenticationWithHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.Nil(t, err)

	psg, err := passage.New("passage", nil)
	require.Nil(t, err)

	t.Run("fail with missing auth token", func(t *testing.T) {
		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})

	t.Run("fail with invalid auth token", func(t *testing.T) {
		req.Header.Add("Bearer", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA")

		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})
}

func TestAuthenticateToken(t *testing.T) {
	psg, err := passage.New(PassageAppID, nil)

	require.Nil(t, err)
	t.Run("valid auth token", func(t *testing.T) {
		_, success := psg.ValidateAuthToken(PassageAuthToken)
		assert.True(t, success)
	})
}

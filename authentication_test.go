package passage_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationWithCookie(t *testing.T) {
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.Nil(t, err)

	psg, err := passage.New(PassageAppID, &passage.Config{
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

	t.Run("valid auth token", func(t *testing.T) {
		// need to create a new request to set a different cookie
		req, err := http.NewRequest("GET", "https://example.com", nil)
		require.Nil(t, err)
		req.AddCookie(&http.Cookie{
			Name:  "psg_auth_token",
			Value: PassageAuthToken,
		})

		userID, err := psg.AuthenticateRequest(req)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, userID)
	})
}

func TestAuthenticationWithHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.Nil(t, err)

	psg, err := passage.New(PassageAppID, &passage.Config{
		HeaderAuth: true,
	})
	require.Nil(t, err)

	t.Run("fail with missing auth token", func(t *testing.T) {
		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})

	t.Run("fail with invalid auth token", func(t *testing.T) {
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA")

		userID, err := psg.AuthenticateRequest(req)
		require.NotNil(t, err)
		assert.Empty(t, userID)
	})

	t.Run("valid auth token", func(t *testing.T) {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", PassageAuthToken))

		userID, err := psg.AuthenticateRequest(req)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, userID)
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

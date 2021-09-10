package passage_test

import (
	"os"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserInfo(t *testing.T) {
	psg, err := passage.New("KZ520QJSiFRLvbBvraaAgYuf", &passage.Config{
		APIKey: os.Getenv("API_KEY"), // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.GetUser("IYQyOzlddrphojERwnMy")
	require.Nil(t, err)
	assert.Equal(t, "IYQyOzlddrphojERwnMy", user.ID)
}
func TestActivateUser(t *testing.T) {
	psg, err := passage.New("KZ520QJSiFRLvbBvraaAgYuf", &passage.Config{
		APIKey: os.Getenv("API_KEY"), // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.ActivateUser("IYQyOzlddrphojERwnMy")
	require.Nil(t, err)
	assert.Equal(t, "IYQyOzlddrphojERwnMy", user.ID)
	assert.Equal(t, true, user.Active)
}
func TestDeactivateUser(t *testing.T) {
	psg, err := passage.New("KZ520QJSiFRLvbBvraaAgYuf", &passage.Config{
		APIKey: os.Getenv("API_KEY"), // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.DeactivateUser("IYQyOzlddrphojERwnMy")
	require.Nil(t, err)
	assert.Equal(t, "IYQyOzlddrphojERwnMy", user.ID)
	assert.Equal(t, false, user.Active)
}

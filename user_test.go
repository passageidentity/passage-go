package passage_test

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// these vars are required as environment variables for testing
var (
	passageAppID  string
	passageApiKey string
	passageUserID string
	randomEmail   = generateRandomEmail(14)
	createdUser   passage.User
)

func generateRandomEmail(prefixLength int) string {
	n := prefixLength
	randomChars := make([]byte, n)
	if _, err := rand.Read(randomChars); err != nil {
		panic(err)
	}
	email := fmt.Sprintf("%X@email.com", randomChars)
	return email
}

func TestMain(t *testing.T) {
	godotenv.Load(".env")

	passageAppID = os.Getenv("PASSAGE_APP_ID")
	passageApiKey = os.Getenv("PASSAGE_API_KEY")
	passageUserID = os.Getenv("PASSAGE_USER_ID")
}

func TestGetUserInfo(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey,
	})
	require.Nil(t, err)

	user, err := psg.GetUser(passageUserID)
	require.Nil(t, err)
	assert.Equal(t, passageUserID, user.ID)
}
func TestActivateUser(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.ActivateUser(passageUserID)
	require.Nil(t, err)
	assert.Equal(t, passageUserID, user.ID)
	assert.Equal(t, passage.StatusActive, user.Status)
}
func TestDeactivateUser(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.DeactivateUser(passageUserID)
	require.Nil(t, err)
	assert.Equal(t, passageUserID, user.ID)
	assert.Equal(t, passage.StatusInactive, user.Status)
}

func TestUpdateUser(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	updateBody := passage.UpdateBody{
		Email: "updatedEmail@123.com",
		Phone: "+15005550006",
	}

	user, err := psg.UpdateUser(passageUserID, updateBody)
	require.Nil(t, err)
	assert.Equal(t, "updatedEmail@123.com", user.Email)
	assert.Equal(t, "+15005550006", user.Phone)
}

func TestCreateUser(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: randomEmail,
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, randomEmail, user.Email)

	createdUser = *user
}

func TestDeleteUser(t *testing.T) {
	psg, err := passage.New(passageAppID, &passage.Config{
		APIKey: passageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	result, err := psg.DeleteUser(createdUser.ID)
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

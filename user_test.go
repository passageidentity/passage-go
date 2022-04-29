package passage_test

import (
	"fmt"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// these vars are required as environment variables for testing
// var (
// 	passageAppID  string
// 	passageApiKey string
// 	passageUserID string
// 	randomEmail   = generateRandomEmail(14)
// 	createdUser   passage.User
// )

// func generateRandomEmail(prefixLength int) string {
// 	n := prefixLength
// 	randomChars := make([]byte, n)
// 	if _, err := rand.Read(randomChars); err != nil {
// 		panic(err)
// 	}
// 	email := fmt.Sprintf("%X@email.com", randomChars)
// 	return email
// }

// func TestMain(t *testing.T) {
// 	godotenv.Load(".env")

// 	passageAppID = os.Getenv("PASSAGE_APP_ID")
// 	passageApiKey = os.Getenv("PASSAGE_API_KEY")
// 	passageUserID = os.Getenv("PASSAGE_USER_ID")
// }

func TestGetUserInfo(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	user, err := psg.GetUser(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
}
func TestActivateUser(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.ActivateUser(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusActive, user.Status)
}
func TestDeactivateUser(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.DeactivateUser(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusInactive, user.Status)
}

func TestUpdateUser(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	updateBody := passage.UpdateBody{
		Email: "updatedEmail@123.com",
		Phone: "+15005550006",
		UserMetadata: map[string]interface{}{
			"example1": "123",
		},
	}
	user, err := psg.UpdateUser(PassageUserID, updateBody)
	require.Nil(t, err)
	assert.Equal(t, "updatedEmail@123.com", user.Email)
	assert.Equal(t, "+15005550006", user.Phone)
	assert.Equal(t, "123", user.UserMetadata["example1"])

	secondUpdateBody := passage.UpdateBody{
		Email: "updatedEmail@123.com",
		Phone: "+15005550006",
		UserMetadata: map[string]interface{}{
			"example1": "456",
		},
	}
	user, err = psg.UpdateUser(PassageUserID, secondUpdateBody)
	require.Nil(t, err)
	assert.Equal(t, "updatedEmail@123.com", user.Email)
	assert.Equal(t, "+15005550006", user.Phone)
	assert.Equal(t, "456", user.UserMetadata["example1"])
}

func TestCreateUser(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	CreatedUser = *user
}

func TestCreateUserWithMetadata(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: fmt.Sprintf("1%v", RandomEmail),
		UserMetadata: map[string]interface{}{
			"example1": "test",
		},
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, "1"+RandomEmail, user.Email)
	assert.Equal(t, "test", user.UserMetadata["example1"].(string))

	CreatedUser = *user
}

func TestDeleteUser(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	result, err := psg.DeleteUser(CreatedUser.ID)
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

func TestListUserDevices(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	devices, err := psg.ListUserDevices(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, 0, len(devices))
}

// NOTE RevokeUserDevice is not tested because it is impossible to spoof webauthn to create a device to then revoke

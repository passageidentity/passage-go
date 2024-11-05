package passage_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInfo(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	user, err := psg.User.Get(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
}

func TestGetInfoByIdentifier(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.User.Create(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	userByIdentifier, err := psg.User.GetByIdentifier(RandomEmail)
	require.Nil(t, err)

	userById, err := psg.User.Get(user.ID)
	require.Nil(t, err)

	assert.Equal(t, user.ID, userById.ID)

	assert.Equal(t, userById, userByIdentifier)
}

func TestGetInfoByIdentifierEmailUpperCase(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.User.Create(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	userByIdentifier, err := psg.User.GetByIdentifier(strings.ToUpper(RandomEmail))
	require.Nil(t, err)

	assert.Equal(t, user.ID, userByIdentifier.ID)
}

func TestGetInfoByIdentifierPhone(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	phone := "+15005550007"
	createUserBody := passage.CreateUserBody{
		Phone: phone,
	}

	user, err := psg.User.Create(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, phone, user.Phone)

	userByIdentifier, err := psg.User.GetByIdentifier(phone)
	require.Nil(t, err)

	userById, err := psg.User.Get(user.ID)
	require.Nil(t, err)

	assert.Equal(t, user.ID, userById.ID)

	assert.Equal(t, userById, userByIdentifier)
}

func TestGetInfoByIdentifierError(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	_, err = psg.User.GetByIdentifier("error@passage.id")
	require.NotNil(t, err)

	expectedMessage := "passage User with Identifier \"error@passage.id\" does not exist"
	assert.Contains(t, err.Error(), expectedMessage)
}

func TestActivate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.User.Activate(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusActive, user.Status)
}
func TestDeactivate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	user, err := psg.User.Deactivate(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusInactive, user.Status)
}

func TestUpdate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	updateBody := passage.UpdateBody{
		Email: "updatedemail-gosdk@passage.id",
		Phone: "+15005550012",
		UserMetadata: map[string]interface{}{
			"example1": "123",
		},
	}
	user, err := psg.User.Update(PassageUserID, updateBody)
	require.Nil(t, err)
	assert.Equal(t, "updatedemail-gosdk@passage.id", user.Email)
	assert.Equal(t, "+15005550012", user.Phone)
	assert.Equal(t, "123", user.UserMetadata["example1"])

	secondUpdateBody := passage.UpdateBody{
		Email: "updatedemail-gosdk@passage.id",
		Phone: "+15005550012",
		UserMetadata: map[string]interface{}{
			"example1": "456",
		},
	}
	user, err = psg.User.Update(PassageUserID, secondUpdateBody)
	require.Nil(t, err)
	assert.Equal(t, "updatedemail-gosdk@passage.id", user.Email)
	assert.Equal(t, "+15005550012", user.Phone)
	assert.Equal(t, "456", user.UserMetadata["example1"])
}

func TestCreate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.User.Create(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	CreatedUser = *user
}

func TestCreateWithMetadata(t *testing.T) {
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

	user, err := psg.User.Create(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, "1"+RandomEmail, user.Email)
	assert.Equal(t, "test", user.UserMetadata["example1"].(string))

	CreatedUser = *user
}

func TestDelete(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	result, err := psg.User.Delete(CreatedUser.ID)
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

func TestListDevices(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	devices, err := psg.User.ListDevices(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, 2, len(devices))
}

// NOTE RevokeUserDevice is not tested because it is impossible to spoof webauthn to create a device to then revoke

func TestSignOut(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	result, err := psg.User.SignOut(PassageUserID)
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

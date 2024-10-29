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

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	user, err := appUser.Get()
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
}

func TestNewAppUserByIdentifier(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	appUser, err := passage.NewAppUserByIdentifier(psg.ID, RandomEmail, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)
	assert.Equal(t, user.ID, appUser.UserID)
}

func TestNewAppUserByIdentifierEmailUpperCase(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	createUserBody := passage.CreateUserBody{
		Email: RandomEmail,
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, RandomEmail, user.Email)

	appUser, err := passage.NewAppUserByIdentifier(psg.ID, strings.ToUpper(RandomEmail), &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)
	assert.Equal(t, user.ID, appUser.UserID)
}

func TestNewAppUserByIdentifierPhone(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	phone := "+15005550007"
	createUserBody := passage.CreateUserBody{
		Phone: phone,
	}

	user, err := psg.CreateUser(createUserBody)
	require.Nil(t, err)
	assert.Equal(t, phone, user.Phone)

	appUser, err := passage.NewAppUserByIdentifier(psg.ID, phone, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)
	assert.Equal(t, user.ID, appUser.UserID)
}

func TestNewAppUserByIdentifierError(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	
	_, err = passage.NewAppUserByIdentifier(psg.ID, "error@passage.id", &passage.Config{
		APIKey: PassageApiKey,
	})
	require.NotNil(t, err)

	expectedMessage := "passage User with Identifier \"error@passage.id\" does not exist"
	assert.Contains(t, err.Error(), expectedMessage)
}

func TestActivate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	user, err := appUser.Activate()
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusActive, user.Status)
}
func TestDeactivate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	user, err := appUser.Deactivate()
	require.Nil(t, err)
	assert.Equal(t, PassageUserID, user.ID)
	assert.Equal(t, passage.StatusInactive, user.Status)
}

func TestUpdate(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	updateBody := passage.UpdateBody{
		Email: "updatedemail-gosdk@passage.id",
		Phone: "+15005550012",
		UserMetadata: map[string]interface{}{
			"example1": "123",
		},
	}
	user, err := appUser.Update(updateBody)
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
	user, err = appUser.Update(secondUpdateBody)
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

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	user, err := appUser.Create(createUserBody)
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

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	user, err := appUser.Create(createUserBody)
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

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	result, err := appUser.Delete()
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

func TestListDevices(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey,
	})
	require.Nil(t, err)

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	devices, err := appUser.ListDevices()
	require.Nil(t, err)
	assert.Equal(t, 2, len(devices))
}

// NOTE RevokeUserDevice is not tested because it is impossible to spoof webauthn to create a device to then revoke

func TestSignOut(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	appUser, err := passage.NewAppUser(psg.ID, PassageUserID, &passage.Config{
		APIKey: PassageApiKey,
	})

	result, err := appUser.SignOut()
	require.Nil(t, err)
	assert.Equal(t, result, true)
}

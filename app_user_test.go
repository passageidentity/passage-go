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
	t.Run("Successful get user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		user, err := psg.User.Get(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.Get(PassageUserID)
		require.NotNil(t, err)
		expectedMessage := "failed to get Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.Get("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}

func TestGetInfoByIdentifier(t *testing.T) {
	t.Run("Success: get user by identifer - exact email", func(t *testing.T) {
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
	})

	t.Run("Success: get user by identifer - email uppercase", func(t *testing.T) {
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
	})

	t.Run("Success: get user by identifer - phone number", func(t *testing.T) {
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
	})

	t.Run("Error: identifier not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.GetByIdentifier("error@passage.id")
		require.NotNil(t, err)

		expectedMessage := "Passage Error - message: passage User with Identifier \"error@passage.id\" does not exist"
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.GetByIdentifier("any@passage.id")
		require.NotNil(t, err)
		expectedMessage := "failed to get Passage User By Identifier"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})
}

func TestActivate(t *testing.T) {
	t.Run("Success: activate user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		user, err := psg.User.Activate(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
		assert.Equal(t, passage.StatusActive, user.Status)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.Activate(PassageUserID)
		require.NotNil(t, err)
		expectedMessage := "failed to activate Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.Activate("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}
func TestDeactivate(t *testing.T) {
	t.Run("Success: deactivate user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		user, err := psg.User.Deactivate(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
		assert.Equal(t, passage.StatusInactive, user.Status)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.Deactivate(PassageUserID)
		require.NotNil(t, err)
		expectedMessage := "failed to deactivate Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.Deactivate("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success: update user", func(t *testing.T) {
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
	})

	t.Run("Error: Bad Request - on phone number and email", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		updateBody := passage.UpdateBody{
			Email: "  ",
			Phone: "  ",
		}
		_, err = psg.User.Update(PassageUserID, updateBody)
		require.NotNil(t, err)
		expectedMessage := "failed to update Passage User's attributes"
		passageBadRequestAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		updateBody := passage.UpdateBody{
			Email: "updatedemail-gosdk@passage.id",
			Phone: "+15005550012",
			UserMetadata: map[string]interface{}{
				"example1": "123",
			},
		}

		_, err = psg.User.Update("PassageUserID", updateBody)
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		updateBody := passage.UpdateBody{
			Email: "updatedemail-gosdk@passage.id",
			Phone: "+15005550012",
			UserMetadata: map[string]interface{}{
				"example1": "123",
			},
		}

		_, err = psg.User.Update(PassageUserID, updateBody)
		require.NotNil(t, err)
		expectedMessage := "failed to update Passage User's attributes"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success: create user", func(t *testing.T) {
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
	})

	t.Run("Success: create user with metadata", func(t *testing.T) {
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
	})

	t.Run("Error: Bad Request - on blank phone number and email", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		createUserBody := passage.CreateUserBody{
			Email: "",
			Phone: "",
		}
		_, err = psg.User.Create(createUserBody)

		require.NotNil(t, err)
		expectedMessage := "failed to create Passage User"
		passageBadRequestAsserts(t, err, expectedMessage)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		createUserBody := passage.CreateUserBody{
			Email: RandomEmail,
		}

		_, err = psg.User.Create(createUserBody)
		require.NotNil(t, err)
		expectedMessage := "failed to create Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success: delete user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		result, err := psg.User.Delete(CreatedUser.ID)
		require.Nil(t, err)
		assert.Equal(t, result, true)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.Delete(CreatedUser.ID)
		require.NotNil(t, err)
		expectedMessage := "failed to delete Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.Delete("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}

func TestListUser(t *testing.T) {
	t.Run("Success: list user devices", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		devices, err := psg.User.ListDevices(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, 2, len(devices))
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.ListDevices(PassageUserID)
		require.NotNil(t, err)
		expectedMessage := "failed to list devices for a Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.ListDevices("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}

// NOTE RevokeUserDevice is not tested because it is impossible to spoof webauthn to create a device to then revoke

func TestRevokeRefreshTokens(t *testing.T) {
	t.Run("Success: sign out user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
		})
		require.Nil(t, err)

		result, err := psg.User.RevokeRefreshTokens(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, result, true)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: "PassageApiKey",
		})
		require.Nil(t, err)

		_, err = psg.User.RevokeRefreshTokens(PassageUserID)
		require.NotNil(t, err)
		expectedMessage := "failed to revoke all refresh tokens for a Passage User"
		passageUnauthorizedAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, &passage.Config{
			APIKey: PassageApiKey,
		})
		require.Nil(t, err)

		_, err = psg.User.RevokeRefreshTokens("PassageUserID")
		require.NotNil(t, err)

		expectedMessage := fmt.Sprintf("Passage Error - message: "+passage.UserIDDoesNotExist, "PassageUserID")
		passageUserNotFoundAsserts(t, err, expectedMessage)
	})
}

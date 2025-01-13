package passage_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/passageidentity/passage-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInfoX(t *testing.T) {
	t.Run("Successful get user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		user, err := psg.User.Get(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		_, err = psg.User.Get(PassageUserID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		_, err = psg.User.Get("PassageUserID")
		require.NotNil(t, err)
		passageUserNotFoundAsserts(t, err)
	})
}

func TestGetInfoByIdentifier(t *testing.T) {
	t.Run("Success: get user by identifer - exact email", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
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
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
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
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		phone := "+15005550007"
		createUserBody := passage.CreateUserArgs{
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
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		_, err = psg.User.GetByIdentifier("error@passage.id")
		require.NotNil(t, err)
		passageCouldNotFindUserByIdentifierAsserts(t, err)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		_, err = psg.User.GetByIdentifier("any@passage.id")
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})
}

func TestActivate(t *testing.T) {
	t.Run("Success: activate user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		user, err := psg.User.Activate(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
		assert.Equal(t, passage.StatusActive, user.Status)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		_, err = psg.User.Activate(PassageUserID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		_, err = psg.User.Activate("PassageUserID")
		require.NotNil(t, err)
		passageUserNotFoundAsserts(t, err)
	})
}
func TestDeactivate(t *testing.T) {
	t.Run("Success: deactivate user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		user, err := psg.User.Deactivate(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, PassageUserID, user.ID)
		assert.Equal(t, passage.StatusInactive, user.Status)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		_, err = psg.User.Deactivate(PassageUserID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		_, err = psg.User.Deactivate("PassageUserID")
		require.NotNil(t, err)

		passageUserNotFoundAsserts(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success: update user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		updateBody := passage.UpdateUserOptions{
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

		secondUpdateBody := passage.UpdateUserOptions{
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

	t.Run("Error: Bad Request: on phone number", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		updateBody := passage.UpdateUserOptions{
			Phone: "  ",
		}
		_, err = psg.User.Update(PassageUserID, updateBody)
		require.NotNil(t, err)
		expectedMessage := "identifier: must be a valid E164 number."
		passageBadRequestAsserts(t, err, expectedMessage)
	})

	t.Run("Error: Bad Request: on email", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		updateBody := passage.UpdateUserOptions{
			Email: "  ",
		}
		_, err = psg.User.Update(PassageUserID, updateBody)
		require.NotNil(t, err)
		expectedMessage := "identifier: must be a valid email address."
		passageBadRequestAsserts(t, err, expectedMessage)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		updateBody := passage.UpdateUserOptions{
			Email: "updatedemail-gosdk@passage.id",
			Phone: "+15005550012",
			UserMetadata: map[string]interface{}{
				"example1": "123",
			},
		}

		_, err = psg.User.Update("PassageUserID", updateBody)
		require.NotNil(t, err)
		passageUserNotFoundAsserts(t, err)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		updateBody := passage.UpdateUserOptions{
			Email: "updatedemail-gosdk@passage.id",
			Phone: "+15005550012",
			UserMetadata: map[string]interface{}{
				"example1": "123",
			},
		}

		_, err = psg.User.Update(PassageUserID, updateBody)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success: create user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
			Email: RandomEmail,
		}

		user, err := psg.User.Create(createUserBody)
		require.Nil(t, err)
		assert.Equal(t, RandomEmail, user.Email)

		CreatedUser = *user
	})

	t.Run("Success: create user with metadata", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
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

	t.Run("Error: Bad Request: on blank phone number and email", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
			Email: "",
			Phone: "",
		}
		_, err = psg.User.Create(createUserBody)

		require.NotNil(t, err)
		assert.Equal(t, "At least one of args.Email or args.Phone is required.", err.Error())
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		createUserBody := passage.CreateUserArgs{
			Email: RandomEmail,
		}

		_, err = psg.User.Create(createUserBody)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success: delete user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		err = psg.User.Delete(CreatedUser.ID)
		require.Nil(t, err)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		err = psg.User.Delete(CreatedUser.ID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		err = psg.User.Delete("PassageUserID")
		require.NotNil(t, err)

		passageUserNotFoundAsserts(t, err)
	})
}

func TestListUser(t *testing.T) {
	t.Run("Success: list user devices", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		devices, err := psg.User.ListDevices(PassageUserID)
		require.Nil(t, err)
		assert.Equal(t, 2, len(devices))
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		_, err = psg.User.ListDevices(PassageUserID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		_, err = psg.User.ListDevices("PassageUserID")
		require.NotNil(t, err)

		passageUserNotFoundAsserts(t, err)
	})
}

func TestRevokeRefreshTokens(t *testing.T) {
	t.Run("Success: sign out user", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		err = psg.User.RevokeRefreshTokens(PassageUserID)
		require.Nil(t, err)
	})

	t.Run("Error: unauthorized", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, "invalid api key")
		require.Nil(t, err)

		err = psg.User.RevokeRefreshTokens(PassageUserID)
		require.NotNil(t, err)
		passageUnauthorizedAsserts(t, err)
	})

	t.Run("Error: not found", func(t *testing.T) {
		psg, err := passage.New(PassageAppID, PassageApiKey)
		require.Nil(t, err)

		err = psg.User.RevokeRefreshTokens("PassageUserID")
		require.NotNil(t, err)

		passageUserNotFoundAsserts(t, err)
	})
}

package passage

import (
	"errors"
)

// GetUser gets a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.Get` instead.
func (a *App) GetUser(userID string) (*User, error) {
	user, err := a.User.Get(userID)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// GetUserByIdentifier gets a user using their identifier
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.GetByIdentifier` instead.
func (a *App) GetUserByIdentifier(identifier string) (*User, error) {
	user, err := a.User.GetByIdentifier(identifier)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// ActivateUser activates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.Activate` instead.
func (a *App) ActivateUser(userID string) (*User, error) {
	user, err := a.User.Activate(userID)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// DeactivateUser deactivates a user using their userID
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.Deactivate` instead.
func (a *App) DeactivateUser(userID string) (*User, error) {
	user, err := a.User.Deactivate(userID)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// UpdateUser receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.Update` instead.
func (a *App) UpdateUser(userID string, updateBody UpdateBody) (*User, error) {
	user, err := a.User.Update(userID, updateBody)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// DeleteUser receives a userID (string), and deletes the corresponding user
// returns true on success, false and error on failure (bool, err)
//
// Deprecated: Use `Passage.User.Delete` instead.
func (a *App) DeleteUser(userID string) (bool, error) {
	if err := a.User.Delete(userID); err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return false, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}

		return false, err
	}

	return true, nil
}

// CreateUser receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
//
// Deprecated: Use `Passage.User.Create` instead.
func (a *App) CreateUser(createUserBody CreateUserBody) (*User, error) {
	user, err := a.User.Create(createUserBody)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return user, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return user, err
}

// ListUserDevices lists a user's devices
// returns a list of devices on success, error on failure
//
// Deprecated: Use `Passage.User.ListDevices` instead.
func (a *App) ListUserDevices(userID string) ([]WebAuthnDevices, error) {
	devices, err := a.User.ListDevices(userID)
	if err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return devices, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}
	}

	return devices, err
}

// RevokeUserDevice gets a user using their userID
// returns a true success, error on failure
//
// Deprecated: Use `Passage.User.RevokeDevice` instead.
func (a *App) RevokeUserDevice(userID, deviceID string) (bool, error) {
	if err := a.User.RevokeDevice(userID, deviceID); err != nil {
		var passageError PassageError
		if errors.As(err, &passageError) {
			return false, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}

		return false, err
	}

	return true, nil
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
//
// Deprecated: Use `Passage.User.RevokeRefreshTokens` instead.
func (a *App) SignOut(userID string) (bool, error) {
	if err := a.User.RevokeRefreshTokens(userID); err != nil {

		var passageError PassageError
		if errors.As(err, &passageError) {
			return false, Error{
				ErrorText:  passageError.Message,
				ErrorCode:  passageError.ErrorCode,
				Message:    passageError.Message,
				StatusCode: passageError.StatusCode,
			}
		}

		return false, err
	}

	return true, nil
}

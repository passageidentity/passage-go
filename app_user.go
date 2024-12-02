package passage

import "errors"

type PassageUser = User
type appUser struct {
	app App
}

func newAppUser(app App) *appUser {
	appUser := appUser{
		app: app,
	}

	return &appUser
}

// Get gets a user using their userID
// returns user on success, error on failure
func (a *appUser) Get(userID string) (*PassageUser, error) {
	user, err := a.app.GetUser(userID)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// GetByIdentifier gets a user using their identifier
// returns user on success, error on failure
func (a *appUser) GetByIdentifier(identifier string) (*PassageUser, error) {
	user, err := a.app.GetUserByIdentifier(identifier)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// Activate activates a user using their userID
// returns user on success, error on failure
func (a *appUser) Activate(userID string) (*PassageUser, error) {
	user, err := a.app.ActivateUser(userID)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// Deactivate deactivates a user using their userID
// returns user on success, error on failure
func (a *appUser) Deactivate(userID string) (*PassageUser, error) {
	user, err := a.app.DeactivateUser(userID)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// Update receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *appUser) Update(userID string, updateBody UpdateBody) (*PassageUser, error) {
	user, err := a.app.UpdateUser(userID, updateBody)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// Delete deletes a user by their user string
// returns error on failure
func (a *appUser) Delete(userID string) error {
	if _, err := a.app.DeleteUser(userID); err != nil {
		var e Error
		if errors.As(err, &e) {
			return PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return err
	}

	return nil
}

// Create receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *appUser) Create(createUserBody CreateUserBody) (*PassageUser, error) {
	user, err := a.app.CreateUser(createUserBody)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return user, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return user, err
	}

	return user, nil
}

// ListDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *appUser) ListDevices(userID string) ([]WebAuthnDevices, error) {
	devices, err := a.app.ListUserDevices(userID)
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return devices, PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return devices, err
	}

	return devices, nil
}

// RevokeDevice gets a user using their userID
// returns error on failure
func (a *appUser) RevokeDevice(userID, deviceID string) error {
	if _, err := a.app.RevokeUserDevice(userID, deviceID); err != nil {
		var e Error
		if errors.As(err, &e) {
			return PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return err
	}

	return nil
}

// RevokeRefreshTokens revokes a users refresh tokens
// returns error on failure
func (a *appUser) RevokeRefreshTokens(userID string) error {
	if _, err := a.app.SignOut(userID); err != nil {
		var e Error
		if errors.As(err, &e) {
			return PassageError{
				Message:    e.Message,
				StatusCode: e.StatusCode,
				ErrorCode:  e.ErrorCode,
			}
		}

		return err
	}

	return nil
}

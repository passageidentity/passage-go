package passage

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
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// GetByIdentifier gets a user using their identifier
// returns user on success, error on failure
func (a *appUser) GetByIdentifier(identifier string) (*PassageUser, error) {
	user, err := a.app.GetUserByIdentifier(identifier)
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// Activate activates a user using their userID
// returns user on success, error on failure
func (a *appUser) Activate(userID string) (*PassageUser, error) {
	user, err := a.app.ActivateUser(userID)
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// Deactivate deactivates a user using their userID
// returns user on success, error on failure
func (a *appUser) Deactivate(userID string) (*PassageUser, error) {
	user, err := a.app.DeactivateUser(userID)
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// Update receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *appUser) Update(userID string, updateBody UpdateBody) (*PassageUser, error) {
	user, err := a.app.UpdateUser(userID, updateBody)
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// Delete deletes a user by their user string
// returns true on success, false and error on failure (bool, err)
func (a *appUser) Delete(userID string) (bool, error) {
	ok, err := a.app.DeleteUser(userID)
	//nolint
	if err != nil {
		return ok, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return ok, nil
}

// Create receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *appUser) Create(createUserBody CreateUserBody) (*PassageUser, error) {
	user, err := a.app.CreateUser(createUserBody)
	//nolint
	if err != nil {
		return user, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return user, nil
}

// ListDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *appUser) ListDevices(userID string) ([]WebAuthnDevices, error) {
	devices, err := a.app.ListUserDevices(userID)
	//nolint
	if err != nil {
		return devices, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return devices, nil
}

// RevokeDevice gets a user using their userID
// returns a true success, error on failure
func (a *appUser) RevokeDevice(userID, deviceID string) (bool, error) {
	ok, err := a.app.RevokeUserDevice(userID, deviceID)
	//nolint
	if err != nil {
		return ok, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return ok, nil
}

// RevokeRefreshTokens revokes a users refresh tokens
// returns true on success, error on failure
func (a *appUser) RevokeRefreshTokens(userID string) (bool, error) {
	ok, err := a.app.SignOut(userID)
	//nolint
	if err != nil {
		return ok, PassageError{
			Message:    err.(Error).Message,
			StatusCode: err.(Error).StatusCode,
			ErrorCode:  err.(Error).ErrorCode,
		}
	}

	return ok, nil
}

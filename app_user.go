package passage

type PassageUser = User
type appUser struct {
	app App
}

const (
	UserIDDoesNotExist     string = "passage User with ID \"%v\" does not exist"
	IdentifierDoesNotExist string = "passage User with Identifier \"%v\" does not exist"
)

func newAppUser(app App) *appUser {
	appUser := appUser{
		app: app,
	}

	return &appUser
}

// Get gets a user using their userID
// returns user on success, error on failure
func (a *appUser) Get(userID string) (*PassageUser, error) {
	return a.app.GetUser(userID)
}

// GetByIdentifier gets a user using their identifier
// returns user on success, error on failure
func (a *appUser) GetByIdentifier(identifier string) (*PassageUser, error) {
	return a.app.GetUserByIdentifier(identifier)
}

// Activate activates a user using their userID
// returns user on success, error on failure
func (a *appUser) Activate(userID string) (*PassageUser, error) {
	return a.app.ActivateUser(userID)
}

// Deactivate deactivates a user using their userID
// returns user on success, error on failure
func (a *appUser) Deactivate(userID string) (*PassageUser, error) {
	return a.app.DeactivateUser(userID)
}

// Update receives an UpdateBody struct, updating the corresponding user's attribute(s)
// returns user on success, error on failure
func (a *appUser) Update(userID string, updateBody UpdateBody) (*PassageUser, error) {
	return a.app.UpdateUser(userID, updateBody)
}

// Delete deletes a user by their user string
// returns true on success, false and error on failure (bool, err)
func (a *appUser) Delete(userID string) (bool, error) {
	return a.app.DeleteUser(userID)
}

// Create receives a CreateUserBody struct, creating a user with provided values
// returns user on success, error on failure
func (a *appUser) Create(createUserBody CreateUserBody) (*PassageUser, error) {
	return a.app.CreateUser(createUserBody)
}

// ListDevices lists a user's devices
// returns a list of devices on success, error on failure
func (a *appUser) ListDevices(userID string) ([]WebAuthnDevices, error) {
	return a.app.ListUserDevices(userID)
}

// RevokeDevice gets a user using their userID
// returns a true success, error on failure
func (a *appUser) RevokeDevice(userID, deviceID string) (bool, error) {
	return a.app.RevokeUserDevice(userID, deviceID)
}

// Signout revokes a users refresh tokens
// returns true on success, error on failure
func (a *appUser) SignOut(userID string) (bool, error) {
	return a.app.SignOut(userID)
}

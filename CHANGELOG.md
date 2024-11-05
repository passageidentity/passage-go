# Changelog

All notable changes to this project will be documented in this file.

## [1.11.3] - 2024-11-30

### Added
- deprecated methods
    - Passage.GetUser(userId: string): User
    - Passage.GetUserByIdentifier(identifier: string): User
    - Passage.ActivateUser(userId: string): User
    - Passage.DeactivateUser(userId: string): User
    - Passage.UpdateUser(userId: string, options: UpdateUserArgs): User
    - Passage.CreateUser(args: CreateUserArgs): User
    - Passage.DeleteUser(userId: string): boolean
    - Passage.ListUserDevices(userId: string): WebAuthnDevice[]
    - Passage.RevokeUserDevice(userId: string, deviceId: string): boolean
    - Passage.ValidateAuthToken(authToken: string): string, bool
- added new methods
    - Passage.User.Get(userId: string): User
    - Passage.User.GetByIdentifier(identifier: string): User
    - Passage.User.Activate(userId: string): User
    - Passage.User.Deactivate(userId: string): User
    - Passage.User.Update(userId: string, options: UpdateUserArgs): User
    - Passage.User.Create(args: CreateUserArgs): User
    - Passage.User.Delete(userId: string): boolean
    - Passage.User.ListDevices(userId: string): WebAuthnDevice[]
    - Passage.User.RevokeDevice(userId: string, deviceId: string): boolean
    - Passage.ValidateJwt(authToken: string): string, bool

## [1.11.2] - 2024-10-24

### Added

- chore: LICENSE file added

### Changed

- docs: README updated
- docs: update Passage Docs link
- ci: pin gorename to v0.24.0 to avoid conflict with go runtime 1.20.0
- test: fix test user identifier conflict

## [1.11.1] - 2024-07-29

### Changed

- chore(deps): bump github.com/lestrrat-go/jwx from 1.2.26 to 1.2.29

## [1.11.1] - 2024-10-24

### Added

- LICENSE file added

### Changed

- README updated

## [1.11.0] - 2024-03-21

### Added

- `GetUserByIdentifier` method has been added
- `ListPaginatedUsersItem` model has been added

## [1.9.0] - 2024-01-30

### Added

- `AppleUserSocialConnection` model has been added

### Changed

- `UserEventInfo` has been renamed to `UserRecentEvent`
- Docs have been moved to `/docs`
- `GithubSocialConnection` has been renamed to `GithubUserSocialConnection`
- `GoogleSocialConnection` has been renamed to `GoogleUserSocialConnection`

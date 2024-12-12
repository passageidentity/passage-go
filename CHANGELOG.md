# Changelog

All notable changes to this project will be documented in this file.

## [1.12.0](https://github.com/passageidentity/passage-go/compare/v1.11.2...v1.12.0) (2024-12-12)


### Features

* add new signatures for user and auth functions ([c63c75e](https://github.com/passageidentity/passage-go/commit/c63c75e44705500f2caf45e4a553c0064f3950b8))
* add parameter guards ([#115](https://github.com/passageidentity/passage-go/issues/115)) ([e61d115](https://github.com/passageidentity/passage-go/commit/e61d115f8b0ef2a3624fbcf3f3053d16feac8d24))
* add passage error class ([#105](https://github.com/passageidentity/passage-go/issues/105)) ([f0e7239](https://github.com/passageidentity/passage-go/commit/f0e72390e1e72f84ae3369dbf77dc6361d05a80e))
* adds new fields to the AppInfo and UserInfo structs ([94748b5](https://github.com/passageidentity/passage-go/commit/94748b51882814f94ce15ff2e7a19de8940c2f29))
* jwt audience validation ([#89](https://github.com/passageidentity/passage-go/issues/89)) ([ae2f00d](https://github.com/passageidentity/passage-go/commit/ae2f00d29fd48fe7bc7a908b4a5062189c3b4ea0))
* remove JWK re-fetch logic ([cf887ab](https://github.com/passageidentity/passage-go/commit/cf887abdb0657bbae751bf664f13a56b965f919c))
* reworks the new create magic link func into three separate functions ([#117](https://github.com/passageidentity/passage-go/issues/117)) ([741b260](https://github.com/passageidentity/passage-go/commit/741b260801b6fc127351597c45432c71ebc9786e))


### Bug Fixes

* updates jwx library to use its thread-safe jwks cache ([#88](https://github.com/passageidentity/passage-go/issues/88)) ([b677920](https://github.com/passageidentity/passage-go/commit/b67792097093386c667d940327d859b3dc7e5e32))

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

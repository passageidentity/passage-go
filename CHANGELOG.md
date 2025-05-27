# Changelog

All notable changes to this project will be documented in this file.

## [2.1.1](https://github.com/passageidentity/passage-go/compare/v2.1.0...v2.1.1) (2025-05-27)


### Bug Fixes

* update JWT parsing to latest library versions ([#133](https://github.com/passageidentity/passage-go/issues/133)) ([55ccc95](https://github.com/passageidentity/passage-go/commit/55ccc95101357e861f95afb1dca4e6f94e1fa878))

## [2.1.0](https://github.com/passageidentity/passage-go/compare/v2.0.0...v2.1.0) (2025-01-14)


### Features

* the Auth and User structs are now public ([#128](https://github.com/passageidentity/passage-go/issues/128)) ([d5ef645](https://github.com/passageidentity/passage-go/commit/d5ef6450df0ef7960028f61d225661f626e93f63))

## [2.0.0](https://github.com/passageidentity/passage-go/compare/v1.12.0...v2.0.0) (2025-01-10)


### ⚠ BREAKING CHANGES

* changes various type and struct names
* rename Passage struct and ctor signature and remove old Error class ([#123](https://github.com/passageidentity/passage-go/issues/123))
* remove deprecated code and rename structs and files ([#119](https://github.com/passageidentity/passage-go/issues/119))

### Features

* add parameter guards for Passage constructor ([#125](https://github.com/passageidentity/passage-go/issues/125)) ([85d18d0](https://github.com/passageidentity/passage-go/commit/85d18d0bbaee045358f65aae1dc1b706b4b98a97))
* change type and add parameter guard for language in magic link options ([#124](https://github.com/passageidentity/passage-go/issues/124)) ([afb86ea](https://github.com/passageidentity/passage-go/commit/afb86eaad3ee2a03efd5420f179143b540739020))
* changes various type and struct names ([d82a3e7](https://github.com/passageidentity/passage-go/commit/d82a3e7a5edc37225941f7ba277e03b0f756cd76))
* remove deprecated code and rename structs and files ([#119](https://github.com/passageidentity/passage-go/issues/119)) ([6445861](https://github.com/passageidentity/passage-go/commit/64458618c3c003e1b89e62a8c7a3337ccfaaaab0))
* rename Passage struct and ctor signature and remove old Error class ([#123](https://github.com/passageidentity/passage-go/issues/123)) ([e4fe6f0](https://github.com/passageidentity/passage-go/commit/e4fe6f0cba28f0010e8e22a13e93f82b3fda7d86))

## [1.12.0](https://github.com/passageidentity/passage-go/compare/v1.11.2...v1.12.0) (2024-12-12)


### Features

* add new signatures for user and auth functions ([c63c75e](https://github.com/passageidentity/passage-go/commit/c63c75e44705500f2caf45e4a553c0064f3950b8))
* add parameter guards ([#115](https://github.com/passageidentity/passage-go/issues/115)) ([e61d115](https://github.com/passageidentity/passage-go/commit/e61d115f8b0ef2a3624fbcf3f3053d16feac8d24))
* add passage error class ([#105](https://github.com/passageidentity/passage-go/issues/105)) ([f0e7239](https://github.com/passageidentity/passage-go/commit/f0e72390e1e72f84ae3369dbf77dc6361d05a80e))
* adds new fields to the AppInfo and UserInfo structs ([94748b5](https://github.com/passageidentity/passage-go/commit/94748b51882814f94ce15ff2e7a19de8940c2f29))
* deps update ([4448a3e](https://github.com/passageidentity/passage-go/commit/4448a3e99b00c810d3dea39900515991ef5b84b3))
* go version from 1.16 --&gt; 1.21 ([4448a3e](https://github.com/passageidentity/passage-go/commit/4448a3e99b00c810d3dea39900515991ef5b84b3))
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

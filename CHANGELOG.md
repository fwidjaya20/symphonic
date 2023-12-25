# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.8] - 2023-12-25

### Added

- Adding `Offset` on RunEvent configuration.

## [1.1.7] - 2023-12-23

### Added

- Support `Kafka` as Event Stream.

### Changed

- Using custom parser instead of cron.WithSeconds().
- Change all native Event Stream engine to Watermill Event Stream.

## [1.1.6] - 2023-12-16

### Changed

- Allow to publish an event without any listeners attached.

## [1.1.5] - 2023-12-09

### Added

- Support `log.Logger` contract on Echo Framework.

### Fixed

- Fix invalid Schedule Timing caused by `cron.WithSecond()` by adding "0" in front of expression.

## [1.1.4] - 2023-12-01

### Fixed

- Unlisted `GetArrayString` function on config contract

## [1.1.3] - 2023-11-12

### Added

- Adding `GetArrayString` function on config

## [1.1.2] - 2023-10-17

### Added

- Adding `GetInt`, `GetInt8`, `GetInt16`, `GetInt32`, `GetInt64` function on config.

### Fixed

- Ignore migrations if no migration or seeder files.

## [1.1.1] - 2023-10-04

### Added

- Support `Queue` feature with `RabbitMQ` driver.

## [1.1.0] - 2023-09-26

### Added

- Support `Queue` feature with `Sync` and `Redis` drivers.

## [1.0.0] - 2023-09-16

### Added

- Support `Config` feature.
- Support `Console` feature.
- Support `Database` feature.
- Support `Event` feature.
- Support `Log` feature.
- Support `Schedule` feature.
- Support `Queue` feature

[1.1.8]: https://github.com/fwidjaya20/symphonic/compare/v1.1.6...v1.1.8
[1.1.7]: https://github.com/fwidjaya20/symphonic/compare/v1.1.6...v1.1.7
[1.1.6]: https://github.com/fwidjaya20/symphonic/compare/v1.1.5...v1.1.6
[1.1.5]: https://github.com/fwidjaya20/symphonic/compare/v1.1.4...v1.1.5
[1.1.4]: https://github.com/fwidjaya20/symphonic/compare/v1.1.3...v1.1.4
[1.1.3]: https://github.com/fwidjaya20/symphonic/compare/v1.1.2...v1.1.3
[1.1.2]: https://github.com/fwidjaya20/symphonic/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/fwidjaya20/symphonic/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/fwidjaya20/symphonic/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/fwidjaya20/symphonic/releases/tag/v1.0.0

# maketen-go

[![CI Status](https://github.com/itchyny/maketen-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/itchyny/maketen-go/actions/workflows/ci.yaml?query=branch:main)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchyny/maketen-go/blob/main/LICENSE)
[![release](https://img.shields.io/github/release/itchyny/maketen-go/all.svg)](https://github.com/itchyny/maketen-go/releases)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/itchyny/maketen-go)](https://pkg.go.dev/github.com/itchyny/maketen-go)

Create 10 from numbers!

## Usage

```text
 $ maketen 1 2 3 4
1 * (2 * 3 + 4)
1 + 2 + 3 + 4
1 * 2 * 3 + 4
 $ maketen 3 7 4 8
(3 - 7 / 4) * 8
 $ maketen 1 1 9 9
(1 + 1 / 9) * 9
```

## Installation

### Homebrew

```sh
brew install itchyny/tap/maketen
```

### Build from source

```sh
go install github.com/itchyny/maketen-go/cmd/maketen@latest
```

## Bug Tracker

Report bug at [Issues - itchyny/maketen-go - GitHub](https://github.com/itchyny/maketen-go/issues).

## Author

itchyny (<https://github.com/itchyny>)

## License

This software is released under the MIT License, see LICENSE.

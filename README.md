# maketen-go
[![CI Status](https://github.com/itchyny/maketen-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/itchyny/maketen-go/actions/workflows/ci.yaml?query=branch:main)

### Create 10 from numbers!

## Usage
```
 $ maketen 1 2 3 4
1 + 2 + 3 + 4
1 * 2 * 3 + 4
1 * (2 * 3 + 4)
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
```bash
go install github.com/itchyny/maketen-go/cmd/maketen@latest
```

## Bug Tracker
Report bug at [Issues・itchyny/maketen-go - GitHub](https://github.com/itchyny/maketen-go/issues).

## Author
itchyny (<https://github.com/itchyny>)

## License
This software is released under the MIT License, see LICENSE.

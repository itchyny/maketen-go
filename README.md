# maketen-go [![CI Status](https://github.com/itchyny/maketen-go/workflows/CI/badge.svg)](https://github.com/itchyny/maketen-go/actions)
### Create 10 from four numbers!

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
go get -u github.com/itchyny/maketen-go/cmd/maketen
```

## Bug Tracker
Report bug at [Issues・itchyny/maketen-go - GitHub](https://github.com/itchyny/maketen-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.

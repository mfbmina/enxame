[![codecov](https://codecov.io/gh/mfbmina/enxame/graph/badge.svg?token=NQJ4SCDISO)](https://codecov.io/gh/mfbmina/enxame)
# enxame
`Enxame` is an HTTP load tester and benchmarking utility made in Go.

## Installation

To install `Enxame`, you must download the installation that works for your OS. You can find all releases [here](https://github.com/mfbmina/enxame/releases).

## Usage

The base command for `Enxame` is the command run. You must pass the URI that will be swarmed, like the following:

```sh
./enxame run https://example.com
```

It accepts a couple of flags. If you need to change the default behavior, you can check them using the help:

```sh
./enxame run -h
```

## Contributing

This project is open source and will always be. If you wish to contribute, you're most welcome! You can help with code, reporting issues, or talking to your friends about it.

### Set up your machine

`Enxame` is written in [Go](https://golang.org/).

Prerequisites:
- [Go 1.21+](https://go.dev/doc/install)

Fork the repo, clone it anywhere, and install dependencies:

```sh
git clone git@github.com:mfbmina/enxame.git
cd enxame
go mod tidy
```

You can ensure everything is all right by running the tests:

```sh
go test ./...
```

### Test your changes

You can create a branch for your changes and try to build from the source as you go:

```sh
go build ./...
```

Remember that all changes need to be covered by tests. Ensure that it is all working by running the tests:

```sh
go test ./...
```

Before you commit the changes, we also suggest you format the files if your editor doesn't do it by itself:

```sh
go fmt ./...
```

### Pushing code

Commit messages should be clear and concise. We follow the convention created by [Conventional Commits](https://www.conventionalcommits.org).

Push your branch to your fork and open a pull request against the main branch.

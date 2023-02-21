# Workflows Enable Inactive

[![Coverage Status][badge-coverage]][coverage]
[![Scrutinizer][badge-quality]][quality]

This repository enable the inactive workflows when state is disabled_inactivity

## Requirements
1. Go version >= 1.19

## Build binary (linux example)

```shell
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux main.go
```

## Run in local
```shell
go run main.go config.json
```

## Run tests
```shell
go test ./...
```

[quality]: https://scrutinizer-ci.com/g/blacktrue/workflows-enable-inactive/
[coverage]: https://scrutinizer-ci.com/g/blacktrue/workflows-enable-inactive/code-structure/main/code-coverage

[badge-quality]: https://img.shields.io/scrutinizer/g/blacktrue/workflows-enable-inactive/main?style=flat-square
[badge-coverage]: https://img.shields.io/scrutinizer/coverage/g/blacktrue/workflows-enable-inactive/main?style=flat-square
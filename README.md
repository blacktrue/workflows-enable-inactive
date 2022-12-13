# Workflows Enable Inactive

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
![Go](https://github.com/tonyskapunk/yctf/workflows/Go/badge.svg)
![Image](https://github.com/tonyskapunk/yctf/workflows/Image/badge.svg)
[![License MIT](https://img.shields.io/github/license/tonyskapunk/yctf?style=plastic)](https://github.com/tonyskapunk/yctf/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/tonyskapunk/yctf?style=plastic)](https://github.com/tonyskapunk/yctf/)
[![Go Report Card](https://goreportcard.com/badge/github.com/tonyskapunk/yctf)](https://goreportcard.com/report/github.com/tonyskapunk/yctf)

# yctf

yocto CTF

> From the Latin/Greek octo (οκτώ), meaning "eight", because it is equal to 1000^−8

Itty bitty teeny weeny Capture the Flag web challenge

The main reason behind this super short CTF is to learn and practice [go](https://go.dev/) using [gin](https://github.com/gin-gonic/gin) framework, also I find pretty fun the CTFs challenges.  This is the result of those two.

## Images

Images are available here:

- quay.io/tonyskapunk/yctf
- docker.pkg.github.com/tonyskapunk/yctf/yctf
- ghcr.io/tonyskapunk/yctf

## Run in a container

### Pull and run

```bash
podman run \
  --name yctf \
  --rm \
  --detach \
  --env GIN_MODE=release \
  --publish 8080:8080 \
  docker.pkg.github.com/tonyskapunk/yctf/yctf:latest

## OR
docker run \
  --name yctf \
  --rm \
  --detach \
  --env GIN_MODE=release \
  --publish 8080:8080 \
  yctf:latest
```

### Build locally

```bash
podman build -t yctf:latest .

## OR

docker build -t yctf:latest .
```

## Run the code

### Build and run

```bash
go build -o yctf .
chmod u+x yctf
./yctf
```

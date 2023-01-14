![Go](https://github.com/tonyskapunk/yctf/workflows/Go/badge.svg)
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
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
- ghcr.io/tonyskapunk/yctf
- docker.pkg.github.com/tonyskapunk/yctf/yctf

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
```

### Build locally

```bash
podman build -t yctf:latest .
```

## Run the code

### Build and run

```bash
go build -o yctf .
chmod u+x yctf
./yctf
```

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/komish"><img src="https://avatars0.githubusercontent.com/u/1837593?v=4" width="100px;" alt=""/><br /><sub><b>Jose R. Gonzalez</b></sub></a><br /><a href="https://github.com/tonyskapunk/yctf/commits?author=komish" title="Documentation">📖</a> <a href="https://github.com/tonyskapunk/yctf/commits?author=komish" title="Tests">⚠️</a> <a href="https://github.com/tonyskapunk/yctf/commits?author=komish" title="Code">💻</a></td>
    <td align="center"><a href="https://tonyskapunk.net"><img src="https://avatars0.githubusercontent.com/u/116447?v=4" width="100px;" alt=""/><br /><sub><b>Tony Garcia</b></sub></a><br /><a href="https://github.com/tonyskapunk/yctf/commits?author=tonyskapunk" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

# golang-cli-template
[![CI](https://github.com/bilalcaliskan/golang-cli-template/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/golang-cli-template/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/golang-cli-template)](https://hub.docker.com/r/bilalcaliskan/golang-cli-template/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/golang-cli-template)](https://goreportcard.com/report/github.com/bilalcaliskan/golang-cli-template)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_golang-cli-template&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_golang-cli-template)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_golang-cli-template&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_golang-cli-template)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_golang-cli-template&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_golang-cli-template)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_golang-cli-template&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_golang-cli-template)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_golang-cli-template&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_golang-cli-template)
[![Release](https://img.shields.io/github/release/bilalcaliskan/golang-cli-template.svg)](https://github.com/bilalcaliskan/golang-cli-template/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/golang-cli-template)](https://github.com/bilalcaliskan/golang-cli-template)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Required Steps
- Single command is mostly enough to prepare project, it will prompt you with some questions about your new project:
  ```shell
  $ make -s prepare-initial-project
  ```

## Additional nice-to-have steps
- If you want to build and publish Docker image:
  - Ensure `DOCKER_USERNAME` has been added as **repository secret on GitHub**
  - Ensure `DOCKER_PASSWORD` has been added as **repository secret on GitHub**
  - Uncomment **line 145** to **line 152** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 32** to **line 50** in [build/package/.goreleaser.yaml](build/package/.goreleaser.yaml)
- If you want to enable https://sonarcloud.io/ integration:
  - Ensure your created repository from that template has been added to https://sonarcloud.io/
  - Ensure `SONAR_TOKEN` has been added as **repository secret** on GitHub
  - Ensure `SONAR_TOKEN` has been added as **dependabot secret** on GitHub
  - Uncomment **line 69** to **line 94** in [.github/workflows/pr.yml](.github/workflows/pr.yml)
  - Uncomment **line 116** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 66** to **line 91** in [.github/workflows/push.yml](.github/workflows/push.yml)
- If you want to create banner:
  - Generate a banner from [here](https://devops.datenkollektiv.de/banner.txt/index.html) and place it inside of [build/ci](build/ci) directory into a file **banner.txt**
  - Uncomment **line 18** and **line 35** to **line 38** in [cmd/root.go](cmd/root.go)
  - Run `go get -u github.com/dimiro1/banner`
- If you want to release as Homebrew Formula:
  - At first, you must have a **formula repository** like https://github.com/bilalcaliskan/homebrew-tap
  - Create an access token on account that has **formula repository** mentioned above item and ensure that token is added as`TAP_GITHUB_TOKEN` **repository secret** on GitHub
  - Uncomment **line 165** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 70** to **line 80** in [build/package/.goreleaser.yaml](build/package/.goreleaser.yaml)
- If you want to mock your interfaces with [mockery](https://github.com/vektra/mockery):
  - Add `generate-mocks` target as a prerequisite to all uncommented targets starting with `test` in [Makefile](Makefile)

## Used Libraries
- [spf13/cobra](https://github.com/spf13/cobra)
- [rs/zerolog](https://github.com/rs/zerolog)

## Development
This project requires below tools while developing:
- [Golang 1.21](https://golang.org/doc/go1.21)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ make pre-commit-setup
```

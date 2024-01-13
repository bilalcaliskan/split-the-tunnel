# split-the-tunnel
[![CI](https://github.com/bilalcaliskan/split-the-tunnel/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/split-the-tunnel/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/split-the-tunnel)](https://goreportcard.com/report/github.com/bilalcaliskan/split-the-tunnel)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_split-the-tunnel&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_split-the-tunnel)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_split-the-tunnel&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_split-the-tunnel)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_split-the-tunnel&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_split-the-tunnel)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_split-the-tunnel&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_split-the-tunnel)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_split-the-tunnel&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_split-the-tunnel)
[![Release](https://img.shields.io/github/release/bilalcaliskan/split-the-tunnel.svg)](https://github.com/bilalcaliskan/split-the-tunnel/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/split-the-tunnel)](https://github.com/bilalcaliskan/split-the-tunnel)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Overview
`split-the-tunnel` is a daemon that runs on Linux hosts. It is designed to add domains that bypass VPN, effectively creating a split tunnel. This allows for more efficient use of network resources and can improve network performance.

## Features
- Runs as a daemon on Linux hosts
- Capable of adding domains to bypass VPN
- Creates a split tunnel for efficient network usage

## Installation
To install `split-the-tunnel`, you can download the latest binary from the [releases page](https://github.com/bilalcaliskan/split-the-tunnel/releases/latest) and add it to your PATH.

## Usage
After installing `split-the-tunnel`, you can start the daemon with the following command:

```shell
$ split-the-tunnel start
$ split-the-tunnel add --domain example.com
```

## Development
This project requires below tools while developing:
- [Golang 1.21](https://golang.org/doc/go1.21)
- [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/) and the rest, simply run below command to prepare your
development environment:
```shell
$ make pre-commit-setup
```

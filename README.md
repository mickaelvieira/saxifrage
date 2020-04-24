# Saxifrage

[![Saxifrage](https://github.com/mickaelvieira/saxifrage/workflows/Saxifrage/badge.svg)](https://github.com/mickaelvieira/saxifrage/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/mickaelvieira/saxifrage)](https://goreportcard.com/report/github.com/mickaelvieira/saxifrage) [![GitHub](https://img.shields.io/github/license/mickaelvieira/saxifrage)](https://github.com/mickaelvieira/saxifrage/blob/master/LICENSE.md)

A CLI tool to manage your SSH keys

## Install

```sh
go get github.com/mickaelvieira/saxifrage
```

You can rename the binary if you prefer using a shorter command. Bear in mind that you will have to rename it after each update.

```sh
mv $(go env GOPATH)/bin/saxifrage $(go env GOPATH)/bin/sax
```

## Update

```sh
go get -u github.com/mickaelvieira/saxifrage
```

## Usage

```sh
$ saxifrage

 NAME:
  saxifrage - A CLI tool to manage your SSH keys

 USAGE:
  saxifrage [command]

 COMMANDS:

  config    Show your SSH configuration
  dump      Dump your SSH configuration
  gen       Generate interactively a SSH key (rsa, dsa, ecdsa, ed25519)
  help      Show this help
```


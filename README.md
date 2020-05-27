# Saxifrage

[![Saxifrage](https://github.com/mickaelvieira/saxifrage/workflows/Saxifrage/badge.svg)](https://github.com/mickaelvieira/saxifrage/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/mickaelvieira/saxifrage)](https://goreportcard.com/report/github.com/mickaelvieira/saxifrage) [![GitHub](https://img.shields.io/github/license/mickaelvieira/saxifrage)](https://github.com/mickaelvieira/saxifrage/blob/master/LICENSE.md)

A CLI tool to manage your SSH keys

## Install

Download the latest version corresponding to your platform from the [releases](https://github.com/mickaelvieira/saxifrage/releases) page

```sh
$ unzip saxifrage-linux-amd64.zip
$ sudo mv sax /usr/local/bin/
$ sudo chmod u+x /usr/local/bin/sax
```

## Upgrade

```sh
$ sax upgrade
```

## Usage

```sh
$ sax

 NAME:
  sax - A CLI tool to manage your SSH keys

 USAGE:
  sax [command]

 COMMANDS:

  dump       Dump your SSH configuration
  gen        Generate interactively a SSH key (rsa, dsa, ecdsa, ed25519)
  help       Show this help
  ls         List SSH configuration sections
  rm         Remove interactively a section and its related SSH keys
  upgrade    Upgrade saxifrage
  version    Display the application version
```

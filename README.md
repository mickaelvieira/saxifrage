# Saxifrage

[![Saxifrage](https://github.com/mickaelvieira/saxifrage/workflows/Saxifrage/badge.svg)](https://github.com/mickaelvieira/saxifrage/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/mickaelvieira/saxifrage)](https://goreportcard.com/report/github.com/mickaelvieira/saxifrage) [![GitHub](https://img.shields.io/github/license/mickaelvieira/saxifrage)](https://github.com/mickaelvieira/saxifrage/blob/stable/LICENSE.md)

A CLI tool to manage your SSH keys

## Install

```sh
$ curl -fsSL https://raw.githubusercontent.com/mickaelvieira/saxifrage/stable/scripts/install | sh
```

## Upgrade

```sh
$ sax upgrade
```

## Bash completion

```sh
$ sax completion > sax.sh
$ sudo mv ./sax.sh /usr/share/bash-completion/completions/sax
$ sudo chown root.root /usr/share/bash-completion/completions/sax
$ sudo chmod 0644 /usr/share/bash-completion/completions/sax
$ source /usr/share/bash-completion/completions/sax
```

## Usage

```sh
$ sax

 NAME:
  Saxifrage - A CLI tool to manage your SSH keys

 USAGE:
  sax [command]

 COMMANDS:

  completion    Generate bash completion
  dump          Dump your SSH configuration
  gen           Generate interactively a SSH key (rsa, dsa, ecdsa, ed25519)
  help          Show this help
  ls            List SSH configuration sections
  rm            Remove interactively a section and its related SSH keys
  upgrade       Upgrade Saxifrage
  version       Display the application version

```
## Running Saxifrage in a container

### Docker container

Build the `docker` image

```sh
$ make docker-build
```

Run the container with `docker`

```sh
$ docker run \
    -v /etc/ssh/ssh_config:/etc/ssh/ssh_config \
    -v ~/.ssh/config:/.ssh/config \
    --tty \
    --name sax \
    --rm \
    saxifrage ls
```

### Buildah Container

Build the `buildah` image

```sh
$ make buildah-build
```

Run the container with `podman`

```sh
$ podman run \
    -v /etc/ssh/ssh_config:/etc/ssh/ssh_config \
    -v ~/.ssh/config:/.ssh/config \
    --tty \
    --name sax \
    --rm \
    saxifrage:latest ls
```

# Home
[![Go Reference](https://pkg.go.dev/badge/github.com/lucasloureiror/AegisPass/cmd/aegis.svg)](https://pkg.go.dev/github.com/lucasloureiror/AegisPass/cmd/aegis)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lucasloureiror/AegisPass?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasloureiror/AegisPass)](https://goreportcard.com/report/github.com/lucasloureiror/AegisPass)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/lucasloureiror/AegisPass)
![GitHub](https://img.shields.io/github/license/lucasloureiror/AegisPass)

AegisPass is a CLI tool designed for Software Engineers and DevOps Engineers/Site Reliability Engineers (SRE) to simplify the process of generating high quality random passwords across various platforms and operating systems. AegisPass uses cryptographically secure random numbers along with random.org API to generate random passwords, providing enhanced security for your applications and services with to generate Randomness Mixing.

AegisPass is written in Go and is available for Linux, Windows, and macOS and it does not use anything out of the standard library and random.org, in the principle of no-trust.

## Installation

To install Aegis you need at least Go v1.18 installed in your machine or you can download the binary from the [releases page](https://www.github.com/lucasloureiror/aegispass/releases).

```bash
go install github.com/lucasloureiror/AegisPass/cmd/aegis
```

## Using AegisPass

Use the help command to get instructions on how to use AegisPass.

```bash
aegis help
```



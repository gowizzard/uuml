# uuml

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/uuml.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/uuml/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/uuml/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/uuml)](https://goreportcard.com/report/github.com/gowizzard/uuml) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/uuml.svg)](https://pkg.go.dev/github.com/gowizzard/uuml) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/uuml)](https://github.com/gowizzard/uuml/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/uuml)](https://github.com/gowizzard/uuml/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/uuml)](https://github.com/gowizzard/uuml/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/uuml)](https://github.com/gowizzard/uuml/blob/master/LICENSE) 

This small library should help you to remove umlauts from your strings. It runs through a list of german umlauts and replaces them with the written out spelling.

## Install

Just run the command in the command line and you will download the package.

```console
go get github.com/gowizzard/uuml
```

## How to use?

You just have to include and call the library. It needs only one string and returns a new string. You can also copy the code from this library and include it directly into your code.

```go
convert := uuml.Convert("Hallöchen, mein Name ist Fränklin Meißter!")
```
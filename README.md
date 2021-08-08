# uuml

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
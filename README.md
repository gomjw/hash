# hash
 

[![GoDoc](https://godoc.org/github.com/gomjw/checkproxy?status.svg)](https://godoc.org/github.com/gomjw/hash)

> Package hash returns the hash strings for strings and files

## Install

```console
go get -u github.com/gomjw/hash/...
```

## Import

```go
import "github.com/gomjw/hash"
```

## Usage

#### func  GetMD5

```go
func GetMD5(text string) string
```
GetMD5 returns the MD5 hash of a string.

#### func  GetMD5FromFile

```go
func GetMD5FromFile(file os.File) string
```
GetMD5FromFile returns the MD5 hash of a file.

#### func  GetSHA1

```go
func GetSHA1(text string) string
```
GetSHA1 returns the SHA1 hash of a string.

#### func  GetSHA1FromFile

```go
func GetSHA1FromFile(file os.File) string
```
GetSHA1FromFile returns the SHA1 hash of a file.

#### func  GetSHA256

```go
func GetSHA256(text string) string
```
GetSHA256 returns the SHA256 hash of a string.

#### func  GetSHA256FromFile

```go
func GetSHA256FromFile(file os.File) string
```
GetSHA256FromFile returns the SHA256 hash of a file.

#### func  GetSHA512

```go
func GetSHA512(text string) string
```
GetSHA512 returns the SHA512 hash of a string.

#### func  GetSHA512FromFile

```go
func GetSHA512FromFile(file os.File) string
```
GetSHA512FromFile returns the SHA512 hash of a file.



---

> Made by the awesome contributors of [@gomjw](https://github.com/gomjw) &nbsp;&middot;&nbsp;
> Maintainer [@MarvinJWendt](https://github.com/MarvinJWendt) &nbsp;&middot;&nbsp;
> Readme template version: 1.1.0

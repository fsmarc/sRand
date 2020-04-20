# sRand

[![GoDoc](https://godoc.org/github.com/fsmarc/sRand/Rand?status.svg)](https://godoc.org/github.com/fsmarc/sRand/Rand)

A pseudo-random string generator written in [golang](http://www.golang.org). It uses a 
charset to define the characters of the generated strings.

## Install
```
go get github.com/fsmarc/sRand/Rand
```

## Usage

Create rand
```go
rnd := rand.New(rand.NewSource(1)) //or use another int64 as Seed.
```

The DefaultCharset will be used
```go
fmt.Prinln(rand.DefaultCharset)
//Output:
//abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
```

Get a random Character from the used Charset
```go
fmt.Println(rnd.Char()) // e.g.: f
```

Get a random string of length 10
```go
fmt.Println(rnd.Stringn(10)) // e.g.: C9wiRt23TZ
```

Use a custom Charset
```go
rnd := rand.NewWithCharset(rand.NewSource(1), "Hello World!")
```

## Examples

See [the project documentation](https://godoc.org/github.com/fsmarc/sRand/Rand) for examples:
* [Usage with DefaultCharset](https://godoc.org/github.com/fsmarc/sRand/Rand#example-New)
* [Usage with CustomCharset](https://godoc.org/github.com/fsmarc/sRand/Rand#example-New)
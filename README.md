# CLIMIT
[![PkgGoDev](https://pkg.go.dev/badge/github.com/thedevop1/climit)](https://pkg.go.dev/github.com/thedevop1/climit) 

CLIMIT is a Go package that provides simple functions to limit concurrency. It has similar API as WaitGroup.

Getting Started
===============

Installing
----------

```sh
$ go get github.com/thedevop1/climit
```

Example
-------
```go
    l := NewLimiter(5)
    for i := 0; i <= 10; i++ {
        l.Get()
        go func() {
            defer l.Done()
            // ...
        }()
    }
    l.Wait()
```
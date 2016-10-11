# Description

This library is a collection of some basic data structures and algorithms in [go](https://golang.org/). For now, the data structures **are not thread-safe**.

# Installation

To install this library, you need to have:
* a [_**basic go installation**_](https://golang.org/dl/) and
* a [_**go workspace**_](https://golang.org/doc/code.html) set up.

For instructions on how to do both, consult [golang.org's Getting Started page](https://golang.org/doc/install) step-by-step. Assuming you have done that, fire up your terminal and go to your workspace. This should be as easy as typing `cd $GOPATH`. Then, type
```
go get github.com/ZeroXLR/goalgo
```
This convenient go command will automatically fetch this library for you. It will even take care of fetching its dependencies.

Once done, you can then use any of this library's data structures/algorithms by importing. For instance, if you need `staque` (a _stack_ and a _queue_ in one structure), just put
```go
import "github.com/ZeroXLR/goalgo/staque"
```
at the top of your go file.
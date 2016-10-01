# Description

This library is a collection of some basic data structures and algorithms in [go](https://golang.org/). For now, the data structures _**are not thread-safe**_. Protect them with [Mutexes](https://golang.org/pkg/sync/#Mutex) within concurrent settings.

# Installation

The only prerequisite for this library is a [basic go installation](https://golang.org/dl/). That said, it will be easiest if you have a _**go workspace**_ set up. For instructions on that, go to [golang.org's Getting Started page](https://golang.org/doc/install).

Assuming your workspace is set up, fire up your terminal and go there. This should be as easy as typing `cd $GOPATH`. Then, type
```
go get github.com/ZeroXLR/goalgo
```
The go command will automatically fetch this github repo for you. Then, you can use any of its data structures/algorithms via importing. For instance, if you need `staque` (a _stack_ and a _queue_ in one structure), put
```go
import "github.com/ZeroXLR/goalgo/staque"
```
at the top of your go file.
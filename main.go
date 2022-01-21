package main

import (
	"example/talk"
	"example/types"
	"strings"
)

var _ types.Manager = &X{}

// A service
// All its public methods will be public endpoints
type X struct {
	// add some helper functions for common services
	talk.Service
}

func main() {
	t.Register(X{})
	t.Register(Y{})
}

// A web page can also call this endpoint with RESTful api, such as
// `curl https://talk.com/x/get`
func (x *X) Get() types.Profile {
	return types.Profile{Name: "Jack"}
}

// Another service
type Y struct {
	talk.Service
}

// Here we call service X inside the service Y, the "talk" lib will generate a special instance of
// types.Manager to proxy the communication between X and Y. It will properly serialize the data
func (y *Y) Format() string {
	// once the go 1.18 released, we can use generics to simplify here
	x := t.Connect(X{}).(types.Manager)

	return strings.ToUpper(x.Get().Name)
}

// mock the implementation
var t talk.Talk = nil

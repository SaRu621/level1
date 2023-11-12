package main

import (
	"fmt"
)

type Target interface {
	request() string
}

type Adaptee struct{}

func (a *Adaptee) specificRequest() string {
	return "specific request"
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) request() string {
	return a.adaptee.specificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{
		adaptee: adaptee,
	}

	result := adapter.request()
	fmt.Println(result)
}

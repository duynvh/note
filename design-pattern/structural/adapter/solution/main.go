package main

import "fmt"

type Handler func() error

func execFunc(f Handler) error {
	return f()
}

type MessageHandler func(message string) error

func handlerAdapter(mdhl MessageHandler, message string) Handler {
	return func() error {
		return mdhl(message)
	}
}

func main() {
	f := func() error {
		fmt.Println("Hello world")
		return nil
	}

	execFunc(f) // OK

	msgHdl := func(message string) error {
		fmt.Println(message)
		return nil
	}

	f = handlerAdapter(msgHdl, "Hello World with Adapter Pattern")
	execFunc(f)
}

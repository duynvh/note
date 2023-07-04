package main

import "fmt"

type Handler func() error

func execFunc(f Handler) error {
	return f()
}

type MessageHandler func(message string) error

func main() {
	f := func() error {
		fmt.Println("Hello World")
		return nil
	}

	execFunc(f)

	// BUT, how I can use MessageHandler with execFunc
	// without modify them
}

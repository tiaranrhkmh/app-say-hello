package main

import (
	"fmt"

	gosayhello "github.com/tiaranrhkmh/go-say-hello"
)

func main() {
	tiara := gosayhello.SayHello("Tiara")
	fmt.Println("Hello", tiara)
}

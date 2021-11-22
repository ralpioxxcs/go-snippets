package main

import (
	"cgocallback/wrapper"
	"fmt"
)

func main() {
	fmt.Println("cgo callback test")

	wrapper.TestWrapper()
}

package main

/*
#include <stdio.h>
void CFoo() {
	printf("hello from C\n");
}
*/
import "C"
import (
	"cgocallback/wrapper"
	"fmt"
)

func main() {
	fmt.Println("cgo callback test")
	C.CFoo()

	wrapper.TestWrapper()
}

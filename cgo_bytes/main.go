package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

char *buf;

void setBuf() {
	printf("setBuf\n");
}

const char* getString() {
	return "test";
}

*/
import "C"
import "fmt"

func call() string {
	cstr := C.getString()
	slice := C.GoString(cstr)
	return slice
}

func main() {

	fmt.Println("go : ", call())

	fmt.Println("vim-go")
}

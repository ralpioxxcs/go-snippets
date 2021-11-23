package wrapper

/*
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include "wrapper.h"
#cgo LDFLAGS: -L../c_lib -lstdc++ -lfoo
#cgo CXXFLAGS: -std=c++14 -I../c_lib

void myCallback(int a);
*/
import "C"
import (
	"fmt"
)

func TestWrapper() {
	C.FooCGO(C.CString("hi"))

	C.DoSomethingCGO(C.callbackFnCGO(C.myCallback))
}

//export myCallback
func myCallback(a C.int) {
	fmt.Printf("callback from go context (%v)\n", int(a))
}

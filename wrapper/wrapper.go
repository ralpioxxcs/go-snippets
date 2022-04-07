package wrapper

/*
#cgo CXXFLAGS: -std=c++14 -I../c_lib
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/../c_lib -L${SRCDIR}/../c_lib -lstdc++ -lfoo

#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include "wrapper.h"

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

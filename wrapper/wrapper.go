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

var callbackFn func(int)

func TestWrapper() {
	C.FooCGO(C.CString("hi"))

	//C.DoSomethingCGO(C.callbackFnCGO(C.myCallback))
}

func SetCallbackFn(cb func(int)) {
	callbackFn = cb

	C.DoSomethingCGO(C.callbackFnCGO(C.myCallback))
}

//export myCallback
func myCallback(a C.int) {
	callbackFn(int(a))
}

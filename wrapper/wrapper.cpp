#include "wrapper.h"

#include "foo.h"

#include <iostream>

void FooCGO(char* str) {
    Foo(str);
}

void DoSomethingCGO(callbackFnCGO fn) {
    DoSomething(static_cast<callbackFn>(fn));
}
#include "wrapper.h"

#include "foo.h"

#include <iostream>

void DoSomethingCGO(int a) {
    DoSomething(a);
}

void RegisterCallback(callbackFnCGO fn) {
    RegisterCallback(static_cast<callbackFn>(fn));
}
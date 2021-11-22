#include "foo.h"

#include <iostream>

callbackFn funcPtr;

void DoSomething(int a) {
	std::cout << "foo::DoSomething()" << std::endl;

	funcPtr(a);
}

void RegisterCallback(callbackFn fn) {
	if(fn != nullptr) {
		funcPtr = std::bind(fn, std::placeholders::_1);
	}
}
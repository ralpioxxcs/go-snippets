#include "foo.h"

#include <iostream>

callbackFn funcPtr;

void DoSomething(callbackFn fn) {
	std::cout << "foo::DoSomething()" << std::endl;
	fn(7);
}

void Foo(std::string str) {
	std::cout << str << std::endl;
}
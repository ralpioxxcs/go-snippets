#ifndef FOO_H
#define FOO_H

#include <functional>
#include <string>

using callbackFn = std::function<void(int a)>;

void DoSomething(callbackFn fn);

void Foo(std::string str);

#endif

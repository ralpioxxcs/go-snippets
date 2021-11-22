#ifndef FOO_H
#define FOO_H

#include <functional>
#include <string>

using callbackFn = std::function<void(int a)>;

void DoSomething(int a);

void RegisterCallback(callbackFn fn);

#endif

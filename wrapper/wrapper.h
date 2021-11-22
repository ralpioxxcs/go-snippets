#ifdef __cplusplus
extern "C" {
#endif

// different to callback typename in "foo.h"
typedef void(*callbackFnCGO)(int a);

void DoSomethingCGO(int a);
void RegisterCallback(callbackFnCGO fn);

#ifdef __cplusplus
}
#endif
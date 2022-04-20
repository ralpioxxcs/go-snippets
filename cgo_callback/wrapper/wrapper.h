#ifdef __cplusplus
extern "C" {
#endif

// different to callback typename in "foo.h"
typedef void(*callbackFnCGO)(int a);

void FooCGO(char* str);

void DoSomethingCGO(callbackFnCGO fn);

#ifdef __cplusplus
}
#endif
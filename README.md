# cgo_callback_example

```
$ pushd c_lib
$ clang++ -shared -std=c++14 foo.cpp  -o libfoo.so
$ popd
$ go build main.go
```


# cgo_callback

## How to run
```
$ g++ -shared -fPIC -std=c++14 c_lib/foo.cpp  -o c_lib/libfoo.so
$ go run main.go
```


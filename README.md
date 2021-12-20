# Generic Std
generic wrap for standard lib of golang.

Generic will be supported in go 1.18. 
But for now, standard lib will not support generic containers.
This lib wraps the standard lib. And add some useful generic functions or containers not implemented in standard lib.
The APIs between the lib and standard lib are not totally compatible. But I try to keep the different minimum.

See Also: [Stream API for Go 1.18](https://github.com/TobiasYin/functional)

## Supported List
### Wrapper for Standard lib
container/heap.Heap
container/list.List
container/ring.Ring

sort.Search
sort.Slice

sync/atomic.Value
sync.Map
sync.Pool

### Addition utils
container/set.Set

## Example
Without Generic
```go
import (
    "sync"
)

func MapTest(){
	var m sync.Map
	m.Store("hello", "1")
	m.Store(2, 1)
	m.Store("Data", 2)
	data, ok := m.Load("hello") // data: interface{}
	if ok {
		realData := data.(string)
		...
    }
	data2, ok := m.Load("Data") // data: interface{}
	if ok {
		realData := data2.(string) // panic, data2 is int ,can't assert to string
		...
	}
}

```
This code will panic, but when you compile, you won't get any error info. 
If use generic, the thing will change. Type Error will be checked when compiling.

With Generic
```go
import (
    "github.com/TobiasYin/generic_std/sync"
)

func MapTest(){
	var m sync.Map[string, string]
	m.Store("hello", "1")
	// m.Store(2, 1) // compile error
	// m.Store("Data", 2) // compile error
	m.Store("Data", "2")
	data, ok := m.Load("hello") // data: string, no need for type assertion
	if ok {
		// use data
		...
    }
	data2, ok := m.Load("Data") // data: string
	if ok {
		// use data2
		// data2 type is string, this will be sure if code can compile. 
		...
	}
}

```
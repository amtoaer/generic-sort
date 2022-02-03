[中文](README.md) | english

---

This is a sort package based on Go generics, only available for Go 1.18 and above.

As you can see, this project simply implements sort.Interface in batches by using generics. The slice type can be passed to the sort.Sort method after wrapping with `Sortable` type. Slices of the following types (and aliases) are supported:

```go
type comparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~string
}
```

Import in your project:

```bash
go get github.com/amtoaer/generic-sort
```

You can use it like this:

```go
package main

import (
	"fmt"
	"sort"

	gsort "github.com/amtoaer/generic-sort"
)

func main() {
	intSlice := []int{1, 3, 2, 4}
	stringSlice := []string{"h", "e", "l", "l", "o"}
	byteSlice := []byte{'h', 'e', 'l', 'l', 'o'}
	sort.Sort(gsort.Sortable[int](intSlice))
	sort.Sort(gsort.Sortable[string](stringSlice))
	sort.Sort(gsort.Sortable[byte](byteSlice))
	fmt.Println(intSlice)    // [1 2 3 4]
	fmt.Println(stringSlice) // [e h l l o]
	fmt.Println(byteSlice)   // [101 104 108 108 111]
}
```

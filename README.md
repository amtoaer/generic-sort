中文 | [english](README_en.md)

---

这是一个基于 Go 泛型的排序包，仅适用于 Go 1.18 版本及以上。

如你所见，该项目通过使用泛型简单批量地实现了 sort.Interface，各切片类型可经 Sortable 类型包装后传递给 sort.Sort 方法。支持以下类型（及别名）的切片：

```go
type comparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~string
}
```

在项目中引入：

```bash
go get github.com/amtoaer/generic-sort
```

你可以像这样使用它：

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

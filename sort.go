package sort

type comparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~string
}

// Sortable a generics slice which implements sort.Interface
type Sortable[T comparable] []T

func (s Sortable[T]) Len() int {
	return len(s)
}

func (s Sortable[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Sortable[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

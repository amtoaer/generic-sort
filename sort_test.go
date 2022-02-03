package sort

import (
	"reflect"
	"sort"
	"testing"
)

type testCase[T comparable] struct {
	input   []T
	wanted  []T
	reverse bool
}

func TestSortIntSlice(t *testing.T) {
	test := []testCase[int]{
		{[]int{1, 4, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 4, 2, 3}, []int{4, 3, 2, 1}, true},
	}
	for _, item := range test {
		var sortInter sort.Interface = Sortable[int](item.input)
		if item.reverse {
			sortInter = sort.Reverse(sortInter)
		}
		if sort.Sort(sortInter); !reflect.DeepEqual(item.input, item.wanted) {
			t.Errorf("test fail, %v doesn't equal to %v", item.input, item.wanted)
		}
	}
}

func TestSortFloat32Slice(t *testing.T) {
	test := []testCase[float32]{
		{[]float32{1.2, 3.1, 2.3, 5.7}, []float32{1.2, 2.3, 3.1, 5.7}, false},
		{[]float32{1.2, 3.1, 2.3, 5.7}, []float32{5.7, 3.1, 2.3, 1.2}, true},
	}
	for _, item := range test {
		var sortInter sort.Interface = Sortable[float32](item.input)
		if item.reverse {
			sortInter = sort.Reverse(sortInter)
		}
		if sort.Sort(sortInter); !reflect.DeepEqual(item.input, item.wanted) {
			t.Errorf("test fail, %v doesn't equal to %v", item.input, item.wanted)
		}
	}
}

func TestSortStringSlice(t *testing.T) {
	test := []testCase[string]{
		{[]string{"Hello", "world", "!", "你好", "世界"}, []string{"!", "Hello", "world", "世界", "你好"}, false},
		{[]string{"Hello", "world", "!", "你好", "世界"}, []string{"你好", "世界", "world", "Hello", "!"}, true},
	}
	for _, item := range test {
		var sortInter sort.Interface = Sortable[string](item.input)
		if item.reverse {
			sortInter = sort.Reverse(sortInter)
		}
		if sort.Sort(sortInter); !reflect.DeepEqual(item.input, item.wanted) {
			t.Errorf("test fail, %v doesn't equal to %v", item.input, item.wanted)
		}
	}
}

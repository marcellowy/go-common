package tools

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewSlice(t *testing.T) {

}

func TestSlice_Index(t *testing.T) {
	var arr = NewTSSlice[int](0)
	var wg sync.WaitGroup
	for i, _ := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		wg.Add(1)
		go func(i int) {
			arr.Append(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	var i, err = arr.Index(0)
	if err != nil {
		t.Error(err)
		return
	}
	if i != 0 {
		t.Error("value error")
		return
	}
}

func TestTSSlice_Remove(t *testing.T) {
	var arr = NewTSSlice[int](0)
	arr.Append(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// output: [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(arr.Get())

	_ = arr.Remove(0, 1)
	_ = arr.Remove(0, 1)
	_ = arr.Remove(0, 1)
	_ = arr.Remove(0, 1)
	_ = arr.Remove(0, 1)

	// output: [5 6 7 8 9]
	fmt.Println(arr.Get())

	if arr.Len() != 5 {
		t.Error("length err")
		return
	}
}

func TestTSSlice_Insert(t *testing.T) {
	var arr = NewTSSlice[int](0)
	arr.Append(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// output: [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(arr.Get())
	arr.Insert(0, 99, 98)
	// output: [99 98 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(arr.Get())

}

func TestTSSlice_Append(t *testing.T) {
	var arr = NewTSSlice[int](0)
	arr.Append(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// output: [0 1 2 3 4 5 6 7 8 9]
	arr.Append(99, 98)
	// output: [0 1 2 3 4 5 6 7 8 9 99 98]
	fmt.Println(arr.Get())
}

func TestTSSlice_Reverse(t *testing.T) {
	var arr = NewTSSlice[int](0)
	arr.Append(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	arr.Reverse()
	//output: [9 8 7 6 5 4 3 2 1 0]
	fmt.Println(arr.Get())
}

func TestTSSlice_Clone(t *testing.T) {
	var arr = NewTSSlice[int](0)
	arr.Append(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// output: [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(arr.Get())

	newArr := arr.Clone()
	// output: [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(newArr.Get())

	fmt.Println(&newArr)
	fmt.Println(&arr)
}

package main

import (
	"os"
	"fmt"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
	if sz < 16 {
		sz = 16
	}
	var arr = make([]T, sz)
	copy(arr, da.array)
	da.capacity = sz
	da.array = arr
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	if index < 0 || index >= da.count {
		return fmt.Errorf("bad index '%d'", index)
	}
	if da.count == da.capacity {
		da.MakeArray(da.capacity * 2)
	}
	da.array = append(da.array[:index+1], da.array[index:]...)
	da.array[index] = itm
	da.count++
	return nil
}

func (da *DynArray[T]) Remove(index int) error {
	if index < 0 || index >= da.count {
		return fmt.Errorf("bad index '%d'", index)
	}
	da.array = append(da.array[:index], da.array[index+1:]...)
	da.count--
	if da.count < da.capacity/2 {
		da.MakeArray(da.capacity / 2)
	}
	return nil
}

func (da *DynArray[T]) Append(itm T) {
	if da.count == da.capacity {
		da.MakeArray(da.capacity * 2)
	}
	da.array[da.count-1] = itm
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index < 0 || index >= da.count {
		return result, fmt.Errorf("bad index '%d'", index)
	}
	result = da.array[index]
	return result, nil
}

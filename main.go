package main

import "fmt"

func main() {
	storage := MemoryStorage{
		index: make(map[string][]int),
		value: make(map[string][]string),
	}
	storage.Set("foo", "bar", 2)
	storage.Set("foo", "bar1", 5)
	fmt.Println(storage.Get("foo", 0))
	fmt.Println(storage.Get("foo", 2))
	fmt.Println(storage.Get("foo", 3))
	fmt.Println(storage.Get("foo", 5))
}

type MemoryStorage struct {
	index map[string][]int
	value map[string][]string
}

func (m *MemoryStorage) Set(key, value string, ts int) bool {
	var arr []int
	var values []string

	a, ok := m.index[key]
	if !ok {
		arr = make([]int, 0)
		values = make([]string, 0)
	} else {
		arr = a
		values = m.value[key]
	}

	arr = append(arr, ts)
	values = append(values, value)

	m.index[key] = arr
	m.value[key] = values

	return true
}

func getIndex(arr []int, value int) int {
	res := -1
	s := 0
	e := len(arr)
	mid := 0

	for {
		tmp := arr[s:e]
		if len(tmp) == 0 {
			return s - 1
		}

		mid = (s + e) / 2

		if arr[mid] > value {
			e = mid
		} else if arr[mid] <= value {
			s = mid + 1
		}
	}

	return res
}

func (m *MemoryStorage) Get(key string, ts int) (bool, string) {
	indexes, ok := m.index[key]
	if !ok {
		return false, ""
	}

	targetIndex := getIndex(indexes, ts)
	if targetIndex == -1 {
		return false, ""
	}

	values, ok := m.value[key]
	if !ok {
		return false, ""
	}
	targetValue := values[targetIndex]

	return true, targetValue
}

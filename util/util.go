package util

import "cmp"

func Unique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	newSlices := make([]T, 0)
	m1 := make(map[T]byte)
	for _, v := range ss {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
			newSlices = append(newSlices, v)
		}
	}
	return newSlices
}

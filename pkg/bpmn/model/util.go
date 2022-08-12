package model

import (
	"fmt"
	"math/rand"
)

func AddToSliceInMap[T any](m map[string][]T, key string, val T) {
	vals, ok := m[key]
	if !ok {
		vals = make([]T, 0)
	}
	vals = append(vals, val)
	m[key] = vals
}

func MergeMaps[T any](maps ...map[string]T) map[string]T {
	var result = make(map[string]T)
	for _, m := range maps {
		for id, el := range m {
			result[id] = el
		}
	}
	return result
}

func MergeMapsToSlice[T Element](maps ...map[string]T) []Element {
	var c = 0
	for _, m := range maps {
		c += len(m)
	}
	var result = make([]Element, c)
	for _, m := range maps {
		for _, v := range m {
			result = append(result, v)
		}
	}
	return result
}

func NewID() string {
	return fmt.Sprintf("%d", rand.Uint64())
}

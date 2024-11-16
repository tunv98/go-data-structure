package slice

func RemoveHead[T any](items []T, k int) []T {
	if k >= len(items) {
		return make([]T, 0)
	}
	return items[k:]
}

func RemoveTail[T any](items []T, k int) []T {
	if k >= len(items) {
		return make([]T, 0)
	}
	return items[:len(items)-k]
}

func RemoveAt[T any](items []T, i int) []T {
	if i < 0 || i > len(items)-1 {
		return items
	}
	copy(items[i:], items[i+1:]) // [1,2,4,5,5]
	return items[:len(items)-1]  // not include element at the index len(items)-1
}

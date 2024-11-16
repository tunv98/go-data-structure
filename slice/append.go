package slice

func PushFirst[T any](s []T, item T) []T {
	newSlice := make([]T, len(s)+1)
	newSlice[0] = item
	copy(newSlice[1:], s)
	return newSlice
}

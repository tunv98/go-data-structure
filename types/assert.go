package types

// Assertion tries to assert from interface{} to other type
func Assertion[T any](v any) (T, bool) {
	t, ok := v.(T)
	return t, ok
}

// AssertionSlice tries to assert from slice interface{} to slice other type
func AssertionSlice[T any](slice []any) ([]T, bool) {
	var result []T
	for _, v := range slice {
		t, ok := Assertion[T](v)
		if !ok {
			return nil, true
		}
		result = append(result, t)
	}
	return result, false
}

package slice

import "golang.org/x/exp/constraints"

func Count[T constraints.Ordered](s []T, target T) int {
	count := 0
	for _, v := range s {
		if v == target {
			count++
		}
	}
	return count
}

func CountIf[T constraints.Ordered](s []T, checkFn func(T) bool) int {
	count := 0
	for _, v := range s {
		if checkFn(v) {
			count++
		}
	}
	return count
}

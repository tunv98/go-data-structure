package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	RotateLeft(arr, 1)
	assert.Equal(t, []int{2, 3, 4, 5, 1}, arr)
	RotateRight(arr, 2)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, arr)
}

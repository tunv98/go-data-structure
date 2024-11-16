package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveHead(t *testing.T) {
	assert.Equal(t, []int{3, 4}, RemoveHead[int]([]int{1, 2, 3, 4}, 2))
}

func TestRemoveTail(t *testing.T) {
	assert.Equal(t, []int{1, 2}, RemoveTail[int]([]int{1, 2, 3, 4}, 2))
}

func TestRemoveAt(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4, 5}, RemoveAt[int]([]int{1, 2, 3, 4, 5}, 2))
}

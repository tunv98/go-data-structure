package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReserve(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	Reserve[int](arr)
	assert.Equal(t, []int{4, 3, 2, 1}, arr)
}

package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stack(t *testing.T) {
	s := NewStack(1, 2, 3, 4, 5)
	s.Push(6, 7, 8)
	s.Pop()
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, s.data)
}

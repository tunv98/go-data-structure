package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertion(t *testing.T) {
	strSlice := []any{"a", "b", "c"}
	wantStr, ok := AssertionSlice[string](strSlice)
	assert.Equal(t, []string{"a", "b", "c"}, wantStr)
	assert.Equal(t, true, ok)

	numSlice := []any{1, 2, 3}
	wantNum, ok := AssertionSlice[int](numSlice)
	assert.Equal(t, []int{1, 2, 3}, wantNum)
	assert.Equal(t, true, ok)
}

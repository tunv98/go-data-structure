package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Bitmap(t *testing.T) {
	maxUserID := 100
	ub := NewBitmap(maxUserID)
	ub.Add(10)
	ub.Add(11)
	ub.Add(12)
	ub.Add(100)
	assert.Equal(t, true, ub.Contains(10))
	assert.Equal(t, true, ub.Contains(11))
	assert.Equal(t, false, ub.Contains(13))
	assert.Equal(t, true, ub.Contains(100))
}

/*
add: O(n)
contains: O(1)
*/

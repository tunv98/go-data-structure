package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_queue(t *testing.T) {
	var q Queue[int]
	q.Enqueue(1)
	q.Dequeue()
	q.Dequeue()
	assert.Equal(t, 0, q.Length())
}

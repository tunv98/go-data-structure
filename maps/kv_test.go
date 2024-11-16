package maps

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestContainsKey(t *testing.T) {
	type args[K comparable, V any] struct {
		m   map[K]V
		key K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[any, any]{
		{
			name: "contain number key",
			args: args[any, any]{
				m: map[any]any{
					1: 3,
					2: 4,
				},
				key: 1,
			},
			want: true,
		},
		{
			name: "contain string key",
			args: args[any, any]{
				m: map[any]any{
					"a": "b",
					"c": "d",
				},
				key: "a",
			},
			want: true,
		},
		{
			name: "no contain key",
			args: args[any, any]{
				m: map[any]any{
					1: 3,
				},
				key: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsKey(tt.args.m, tt.args.key); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeysValues(t *testing.T) {
	myMap := map[int]int{
		1: 2,
		2: 3,
		3: 4,
		4: 5,
	}
	keys := Keys[int, int](myMap)
	assert.Equal(t, []int{1, 2, 3, 4}, keys)

	values := Values[int, int](myMap)
	sort.Ints(values)
	assert.Equal(t, []int{2, 3, 4, 5}, values)
}

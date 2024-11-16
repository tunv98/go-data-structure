package slice

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func TestCount(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s      []T
		target T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "success",
			args: args[int]{
				s:      []int{1, 3, 3, 4, 5},
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountIf(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s       []T
		checkFn func(T) bool
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "success",
			args: args[int]{
				s: []int{1, 2, 3, 4, 3, 5},
				checkFn: func(i int) bool {
					return i > 3
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIf(tt.args.s, tt.args.checkFn); got != tt.want {
				t.Errorf("CountIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

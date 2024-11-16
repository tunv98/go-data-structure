package slice

import (
	"reflect"
	"testing"
)

func Test_pushFront(t *testing.T) {
	type args[T any] struct {
		s    []T
		item T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "success",
			args: args[int]{
				s:    []int{2, 3, 4, 5},
				item: 1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PushFirst(tt.args.s, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

package fp

import (
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	type args[D any, K comparable, V any] struct {
		data []D
		f    func(D) (K, V)
	}
	type testCase[D any, K comparable, V any] struct {
		name string
		args args[D, K, V]
		want map[K]V
	}
	tests := []testCase[int, int, int]{
		{
			name: "",
			args: args[int, int, int]{
				data: []int{1, 3, 4},
				f: func(i int) (int, int) {
					return i, i * i
				},
			},
			want: map[int]int{
				1: 1,
				3: 9,
				4: 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToMap(tt.args.data, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

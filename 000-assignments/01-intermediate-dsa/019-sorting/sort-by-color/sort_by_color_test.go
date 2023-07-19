package sortbycolor

import (
	"reflect"
	"testing"
)

func TestSortByColor(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{0, 1, 2, 0, 1, 2}}, []int{0, 0, 1, 1, 2, 2}},
		{"test2", args{[]int{0, 1, 2, 0, 1, 2, 0, 1, 2}}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
		{"test3", args{[]int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByColor(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSortByColor-8   	143825518	         8.269 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSortByColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortByColor([]int{0, 1, 2, 0, 1, 2})
	}
}

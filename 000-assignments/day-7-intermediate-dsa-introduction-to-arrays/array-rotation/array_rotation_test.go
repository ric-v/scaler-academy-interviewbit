package arrayrotation

import (
	"reflect"
	"testing"
)

func TestArrayRotation(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}, 1}, []int{2, 3, 4, 5, 6, 1}},
		{"test2", args{[]int{1, 2, 3, 4, 5, 6}, 2}, []int{3, 4, 5, 6, 1, 2}},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6}, 3}, []int{4, 5, 6, 1, 2, 3}},
		{"test4", args{[]int{1, 2, 3, 4, 5, 6}, 4}, []int{5, 6, 1, 2, 3, 4}},
		{"test5", args{[]int{1, 2, 3, 4, 5, 6}, 5}, []int{6, 1, 2, 3, 4, 5}},
		{"test6", args{[]int{1, 2, 3, 4, 5, 6}, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"test7", args{[]int{1, 2, 3, 4}, 2}, []int{3, 4, 1, 2}},
		{"test8", args{[]int{2, 5, 6}, 1}, []int{6, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRotation(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkArrayRotation-8   	13042202	        91.60 ns/op	       0 B/op	       0 allocs/op
func BenchmarkArrayRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArrayRotation([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}, 100)
	}
}

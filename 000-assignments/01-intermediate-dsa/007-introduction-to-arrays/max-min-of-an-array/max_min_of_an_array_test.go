package maxminofanarray

import (
	"testing"
)

func TestMaxMinOfAnArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, 6},
		{"test2", args{[]int{5, 4, 3, 2, 1}}, 6},
		{"test3", args{[]int{1, 1, 1, 1, 1}}, 2},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}}, 9},
		{"test5", args{[]int{-2, 1, -4, 5, 3}}, 1},
		{"test6", args{[]int{1, 3, 4, 1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxMinSumOfAnArray(tt.args.A); got != tt.want {
				t.Errorf("MaxMinOfAnArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMaxMinOfAnArray-8   	15055761	        87.06 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMaxMinOfAnArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxMinSumOfAnArray([]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6})
	}
}

func TestMinMax(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, 1, 5},
		{"test2", args{[]int{5, 4, 3, 2, 1}}, 1, 5},
		{"test3", args{[]int{1, 1, 1, 1, 1}}, 1, 1},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}}, 1, 8},
		{"test5", args{[]int{-2, 1, -4, 5, 3}}, -4, 5},
		{"test6", args{[]int{1, 3, 4, 1}}, 1, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinMax(tt.args.A)
			if got != tt.want {
				t.Errorf("MinMax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinMax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// BenchmarkMinMax-8   	18636838	        65.35 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMinMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinMax([]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6})
	}
}

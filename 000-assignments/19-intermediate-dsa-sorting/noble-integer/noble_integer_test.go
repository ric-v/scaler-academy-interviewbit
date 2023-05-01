package main

import "testing"

func TestNobleInteger(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 2, 1, 3}}, 1},
		{"test2", args{[]int{1, 1, 3, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NobleInteger(tt.args.A); got != tt.want {
				t.Errorf("NobleInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkNobleInteger-8   	 9070053	       114.8 ns/op	      56 B/op	       2 allocs/op
func BenchmarkNobleInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NobleInteger([]int{3, 2, 1, 3})
	}
}

package kthsmallestelement

import "testing"

func TestKthSmallestElement(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 1, 4, 3, 2}, 3}, 2},
		{"test2", args{[]int{7, 10, 4, 3, 20, 15}, 3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthSmallestElement(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("KthSmallestElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKthSmallestElement-8   	44933217	        25.33 ns/op	       0 B/op	       0 allocs/op
func BenchmarkKthSmallestElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KthSmallestElement([]int{2, 1, 4, 3, 2}, 3)
	}
}

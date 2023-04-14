package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 2, 3, 1}}, 3},
		{"test2", args{[]int{1, 2, 2, 3, 1, 3, 5}}, 5},
		{"test3", args{[]int{1, 2, 2, 3, 1, 3, 5, 5, 6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.args.A); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSingleNumber-8   	330586796	         3.546 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumber([]int{1, 2, 2, 3, 1, 3, 5, 5, 6})
	}
}

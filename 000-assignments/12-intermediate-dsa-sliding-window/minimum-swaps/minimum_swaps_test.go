package minimumswaps

import "testing"

func TestMinimumSwaps(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 12, 10, 3, 14, 10, 5}, 8}, 2},
		{"test2", args{[]int{15, 17, 100, 11}, 20}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSwaps(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("MinimumSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMinimumSwaps-8   	84733788	        14.75 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMinimumSwaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinimumSwaps([]int{1, 12, 10, 3, 14, 10, 5}, 8)
	}
}

package goodsubarrays

import "testing"

func TestGoodSubarrays(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, 4}, 6},
		{"test2", args{[]int{13, 16, 16, 15, 9, 16, 2, 7, 6, 17, 3, 9}, 65}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoodSubarrays(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("GoodSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkGoodSubarrays-8   	38478679	        31.00 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGoodSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoodSubarrays([]int{1, 2, 3, 4, 5}, 4)
	}
}

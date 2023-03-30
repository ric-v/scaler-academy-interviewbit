package countingsubarrays

import "testing"

func TestCountingSubarrays(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 5, 6}, 10}, 4},
		{"test2", args{[]int{1, 11, 2, 3, 15}, 10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountingSubarrays(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("CountingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountingSubarrays-8   	142289883	         8.407 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountingSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSubarrays([]int{2, 5, 6}, 10)
	}
}

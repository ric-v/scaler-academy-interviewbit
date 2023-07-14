package longestconsecutivesequence

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{100, 4, 200, 1, 3, 2}}, 4},
		{"test2", args{[]int{2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutiveSequence(tt.args.A); got != tt.want {
				t.Errorf("LongestConsecutiveSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLongestConsecutiveSequence-8   	 6556087	       180.6 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLongestConsecutiveSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestConsecutiveSequence([]int{100, 4, 200, 1, 3, 2})
	}
}

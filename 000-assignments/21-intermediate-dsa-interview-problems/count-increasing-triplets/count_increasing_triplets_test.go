package countincreasingtriplets

import "testing"

func TestCountIncreasingTriplets(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 4, 3}}, 2},
		{"test2", args{[]int{2, 1, 2, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIncreasingTriplets(tt.args.A); got != tt.want {
				t.Errorf("CountIncreasingTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountIncreasingTriplets-8   	71483460	        16.92 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountIncreasingTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountIncreasingTriplets([]int{1, 2, 4, 3})
	}
}

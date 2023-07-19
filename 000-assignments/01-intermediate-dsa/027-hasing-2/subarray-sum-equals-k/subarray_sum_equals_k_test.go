package subarraysumequalsk

import "testing"

func TestSubArraySumEqualsK(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 0, 1}, 1}, 4},
		{"test2", args{[]int{0, 0, 0}, 0}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraySumEqualsK(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SubArraySumEqualsK() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubArraySumEqualsK-8   	14808579	        70.49 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSubArraySumEqualsK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubArraySumEqualsK([]int{1, 0, 1}, 1)
	}
}

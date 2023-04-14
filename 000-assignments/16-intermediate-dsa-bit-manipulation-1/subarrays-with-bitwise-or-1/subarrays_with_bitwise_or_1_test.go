package subarrayswithbitwiseor1

import "testing"

func TestSubArraysWithBitwiseOr(t *testing.T) {
	type args struct {
		A int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{3, []int{1, 0, 1}}, 5},
		{"test2", args{2, []int{1, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraysWithBitwiseOr(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SubArraysWithBitwiseOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubArraysWithBitwiseOr-8   	644031457	         1.868 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSubArraysWithBitwiseOr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubArraysWithBitwiseOr(3, []int{1, 0, 1})
	}
}

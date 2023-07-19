package kthsymbolhard

import "testing"

func TestKthSymbolHard(t *testing.T) {
	type args struct {
		A int
		B int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{3, 0}, 0},
		{"test2", args{4, 4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthSymbolHard(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("KthSymbolHard() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKthSymbolHard-8   	203710386	         6.182 ns/op	       0 B/op	       0 allocs/op
func BenchmarkKthSymbolHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KthSymbolHard(4, 4)
	}
}

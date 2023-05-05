package kthsymbol

import "testing"

func TestKthSymbol(t *testing.T) {
	type args struct {
		A int
		B int
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
			if got := KthSymbol(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("KthSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKthSymbol-8   	27377575	        45.10 ns/op	       0 B/op	       0 allocs/op
func BenchmarkKthSymbol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KthSymbol(3, 0)
	}
}

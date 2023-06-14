package gcd

import "testing"

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{a: 1, b: 1}, 1},
		{"test2", args{a: 24, b: 16}, 8},
		{"test3", args{a: 36, b: 24}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 22021944                53.41 ns/op            0 B/op          0 allocs/op
func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(24, 16)
	}
}

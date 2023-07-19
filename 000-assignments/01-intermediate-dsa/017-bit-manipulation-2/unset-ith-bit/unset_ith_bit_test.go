package unsetithbit

import "testing"

func TestUnsetIthBit(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{A: 4, B: 1}, 4},
		{"test2", args{A: 5, B: 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsetIthBit(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("UnsetIthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkUnsetIthBit-8   	1000000000	         0.2829 ns/op	       0 B/op	       0 allocs/op
func BenchmarkUnsetIthBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnsetIthBit(4, 1)
	}
}

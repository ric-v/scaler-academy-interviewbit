package unsetxbitsfromright

import "testing"

func TestUnsetXBitsFromRight(t *testing.T) {
	type args struct {
		A int64
		B int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{A: 25, B: 3}, 24},
		{"test2", args{A: 37, B: 3}, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsetXBitsFromRight(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("UnsetXBitsFromRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkUnsetXBitsFromRight-8   	1000000000	         0.2775 ns/op	       0 B/op	       0 allocs/op
func BenchmarkUnsetXBitsFromRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnsetXBitsFromRight(25, 3)
	}
}

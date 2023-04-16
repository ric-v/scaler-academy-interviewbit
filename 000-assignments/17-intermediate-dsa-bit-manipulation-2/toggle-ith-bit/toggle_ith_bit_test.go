package toggleithbit

import "testing"

func TestToggleIthBit(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{A: 4, B: 1}, 6},
		{"test2", args{A: 5, B: 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleIthBit(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("ToggleIthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkToggleIthBit-8   	1000000000	         0.2747 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToggleIthBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToggleIthBit(4, 1)
	}
}

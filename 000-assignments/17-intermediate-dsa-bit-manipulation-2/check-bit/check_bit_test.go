package checkbit

import "testing"

func TestCheckBit(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4, 1}, 0},
		{"test2", args{5, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBit(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("CheckBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCheckBit-8   	1000000000	         0.2712 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCheckBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckBit(5, 2)
	}
}

package setbit

import "testing"

func TestSetBit(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{3, 5}, 40},
		{"test2", args{4, 4}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBit(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSetBit-8   	1000000000	         0.2991 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSetBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetBit(3, 5)
	}
}

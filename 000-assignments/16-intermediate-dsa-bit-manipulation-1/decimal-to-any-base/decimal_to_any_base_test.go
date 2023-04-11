package decimaltoanybase

import "testing"

func TestDecimalToAnyBase(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{A: 4, B: 3}, 11},
		{"test2", args{A: 4, B: 2}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalToAnyBase(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("DecimalToAnyBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkDecimalToAnyBase-8   	655057922	         1.815 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDecimalToAnyBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalToAnyBase(4, 3)
	}
}

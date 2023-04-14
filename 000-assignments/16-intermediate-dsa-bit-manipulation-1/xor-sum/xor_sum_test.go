package xorsum

import "testing"

func TestXorSum(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{6, 12}, 10},
		{"test2", args{4, 9}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XorSum(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("XorSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkXorSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XorSum(6, 12)
	}
}

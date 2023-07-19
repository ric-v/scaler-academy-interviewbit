package numberof1bits

import "testing"

func TestNumberOf1Bits(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{11}, 3},
		{"test2", args{128}, 1},
		{"test3", args{6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOf1Bits(tt.args.A); got != tt.want {
				t.Errorf("NumberOf1Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkNumberOf1Bits-8   	361822638	         3.416 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNumberOf1Bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOf1Bits(11)
	}
}

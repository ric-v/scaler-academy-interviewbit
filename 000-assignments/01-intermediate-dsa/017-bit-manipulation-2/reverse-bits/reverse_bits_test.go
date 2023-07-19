package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"test1", args{num: 0}, 0},
		{"test2", args{num: 3}, 3221225472},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBits(tt.args.num); got != tt.want {
				t.Errorf("ReverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkReverseBits-8   	103756369	        11.36 ns/op	       0 B/op	       0 allocs/op
func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBits(3)
	}
}

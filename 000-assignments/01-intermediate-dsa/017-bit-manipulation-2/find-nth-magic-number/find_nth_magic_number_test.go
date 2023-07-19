package findnthmagicnumber

import "testing"

func TestFindNthMagicNumber(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{A: 3}, 30},
		{"test2", args{A: 10}, 650},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNthMagicNumber(tt.args.A); got != tt.want {
				t.Errorf("FindNthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindNthMagicNumber-8   	215500984	         6.331 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFindNthMagicNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindNthMagicNumber(100)
	}
}

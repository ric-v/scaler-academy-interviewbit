package lengthoflongestconsecutiveones

import "testing"

func TestLenOfLongestConsecutiveOnes(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"111000"}, 3},
		{"test2", args{"111011101"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenOfLongestConsecutiveOnes(tt.args.A); got != tt.want {
				t.Errorf("LenOfLongestConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLenOfLongestConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LenOfLongestConsecutiveOnes("111011101")
	}
}

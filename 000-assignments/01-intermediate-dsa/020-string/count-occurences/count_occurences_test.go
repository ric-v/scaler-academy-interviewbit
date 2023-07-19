package countoccurences

import "testing"

func TestCountOccurences(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"bobob"}, 2},
		{"test2", args{"bobobobobobob"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOccurences(tt.args.A); got != tt.want {
				t.Errorf("CountOccurences() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountOccurences-8   	 2120692	       520.3 ns/op	      96 B/op	      12 allocs/op
func BenchmarkCountOccurences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountOccurences("bobobobobobob")
	}
}

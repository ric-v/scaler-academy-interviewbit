package rabinkarp

import "testing"

func TestRabinKarp(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"acbac", "ac"}, 2},
		{"test2", args{"aaaa", "aa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RabinKarp(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("RabinKarp() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkRabinKarp-8   	14240584	        85.36 ns/op	      16 B/op	       1 allocs/op
func BenchmarkRabinKarp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RabinKarp("acbac", "ac")
	}
}

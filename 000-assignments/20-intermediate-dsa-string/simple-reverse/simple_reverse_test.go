package simplereverse

import "testing"

func TestSimpleReverse(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"abc"}, "cba"},
		{"test2", args{"a"}, "a"},
		{"test3", args{"the quick brown fox"}, "xof nworb kciuq eht"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleReverse(tt.args.A); got != tt.want {
				t.Errorf("SimpleReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSimpleReverse-8   	  485624	      2208 ns/op	    1312 B/op	      85 allocs/op
func BenchmarkSimpleReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleReverse("the quick brown fox jumps over the lazy dog")
	}
}

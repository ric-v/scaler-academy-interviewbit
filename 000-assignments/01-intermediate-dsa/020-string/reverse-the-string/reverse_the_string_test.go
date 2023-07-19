package reversethestring

import "testing"

func TestReverseTheString(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"the sky is blue"}, "blue is sky the"},
		{"test2", args{"this is ib"}, "ib is this"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseTheString(tt.args.A); got != tt.want {
				t.Errorf("ReverseTheString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkReverseTheString-8   	 3660481	       330.2 ns/op	     192 B/op	       5 allocs/op
func BenchmarkReverseTheString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseTheString("the sky is blue")
	}
}

package checkanagrams

import "testing"

func TestCheckAnagrams(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"bobob", "bobob"}, 1},
		{"test2", args{"bat", "cat"}, 0},
		{"test3", args{"secure", "rescue"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckAnagrams(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("CheckAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCheckAnagrams-8   	 8555893	       125.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCheckAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckAnagrams("bobob", "bobob")
	}
}

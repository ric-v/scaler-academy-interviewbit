package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{[]string{"abcdefgh", "aefghijk", "abcefgh"}}, "a"},
		{"test2", args{[]string{"abab", "ab", "abcd"}}, "ab"},
		{"test3", args{[]string{"abcdefgh", "abcefghijk", "abcefgh", "abcefgh"}}, "abc"},
		{"test4", args{[]string{"abcd", "abcd", "abcd"}}, "abcd"},
		{"test5", args{[]string{"abcd", "abcde", "efgh"}}, ""},
		{"test5", args{[]string{"abcd", "efgh", "abcde"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.A); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLongestCommonPrefix-8   	18909409	        60.43 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix([]string{"abcdefgh", "aefghijk", "abcefgh"})
	}
}

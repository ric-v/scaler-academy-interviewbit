package longestpalindromicsubstring

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"aaaabaaa"}, "aaabaaa"},
		{"test2", args{"abba"}, "abba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromicSubstring(tt.args.A); got != tt.want {
				t.Errorf("LongestPalindromicSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLongestPalindromicSubstring-8   	12911208	        92.83 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLongestPalindromicSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestPalindromicSubstring("aaaabaaa")
	}
}

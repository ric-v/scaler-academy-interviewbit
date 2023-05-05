package checkpalindrome

import "testing"

func TestCheckPalindrome(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"naman"}, 1},
		{"test2", args{"abba"}, 1},
		{"test3", args{"ab"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPalindrome(tt.args.A); got != tt.want {
				t.Errorf("CheckPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCheckPalindrome-8   	231303309	         5.036 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCheckPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckPalindrome("naman")
	}
}

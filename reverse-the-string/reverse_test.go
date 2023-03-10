package reverse

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				A: "",
			},
			want: "",
		},
		{
			name: "single character",
			args: args{
				A: "a",
			},
			want: "a",
		},
		{
			name: "two words",
			args: args{
				A: " hello    world  ",
			},
			want: "world hello",
		},
		{
			name: "three words with space",
			args: args{
				A: " hello new world  ",
			},
			want: "world new hello",
		},
		{
			name: "four words",
			args: args{
				A: "the sky is blue",
			},
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.A); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkReverse-8   	40837299	        26.85 ns/op	       8 B/op	       1 allocs/op
// BenchmarkReverse-8   	 4697088	        244.3 ns/op	     112 B/op	       5 allocs/op
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("the sky is blue")
	}
}

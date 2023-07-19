package addbinarystrings

import "testing"

func TestAddBinaryStrings(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"100", "11"}, "111"},
		{"test2", args{"110", "10"}, "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinaryStrings(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("AddBinaryStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkAddBinaryStrings-8   	11562508	       105.7 ns/op	      24 B/op	       3 allocs/op
func BenchmarkAddBinaryStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddBinaryStrings("100", "11")
	}
}

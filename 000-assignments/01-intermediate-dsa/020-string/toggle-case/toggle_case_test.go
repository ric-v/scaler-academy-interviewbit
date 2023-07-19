package togglecase

import "testing"

func TestToggleCase(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"abc def"}, "ABC DEF"},
		{"test2", args{"AbC dEf"}, "aBc DeF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleCase(tt.args.A); got != tt.want {
				t.Errorf("ToggleCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkToggleCase-8   	45278104	        25.64 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToggleCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToggleCase("aBc DeF gHi JkL")
	}
}

package doublecharactertrouble

import "testing"

func TestDoubleCharacterTrouble(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"abccbc"}, "ac"},
		{"test2", args{"ac"}, "ac"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleCharacterTrouble(tt.args.s); got != tt.want {
				t.Errorf("DoubleCharacterTrouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkDoubleCharacterTrouble-8   	13729348	        90.07 ns/op	      24 B/op	       2 allocs/op
func BenchmarkDoubleCharacterTrouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoubleCharacterTrouble("abccbc")
	}
}

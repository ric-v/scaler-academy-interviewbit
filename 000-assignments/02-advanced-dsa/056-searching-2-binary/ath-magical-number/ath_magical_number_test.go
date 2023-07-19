package athmagicalnumber

import "testing"

func TestAthMagicalNumber(t *testing.T) {
	type args struct {
		A int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1, 2, 3}, 2},
		{"test2", args{4, 2, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AthMagicalNumber(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("AthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAthMagicalNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AthMagicalNumber(4, 2, 3)
	}
}

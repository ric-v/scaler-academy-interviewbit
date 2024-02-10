package ndigitnumber

import "testing"

func TestNDigitNumber(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NDigitNumber(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("NDigitNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

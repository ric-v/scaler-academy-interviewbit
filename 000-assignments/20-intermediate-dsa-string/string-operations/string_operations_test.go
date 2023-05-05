package stringoperations

import "testing"

func TestStringOperations(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringOperations(tt.args.A); got != tt.want {
				t.Errorf("StringOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

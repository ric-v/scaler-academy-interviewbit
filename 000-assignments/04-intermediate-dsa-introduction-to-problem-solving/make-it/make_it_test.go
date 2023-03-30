package makeit

import "testing"

func TestMakeIt(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test Case 1", args: args{A: 1, B: 2}, want: 2},
		{name: "Test Case 2", args: args{A: 2, B: 3}, want: 4},
		{name: "Test Case 3", args: args{A: 3, B: 4}, want: 6},
		{name: "Test Case 4", args: args{A: 4, B: 5}, want: 8},
		{name: "Test Case 5", args: args{A: 5, B: 6}, want: 10},
		{name: "Test Case 6", args: args{A: 6, B: 7}, want: 12},
		{name: "Test Case 7", args: args{A: 7, B: 8}, want: 14},
		{name: "Test Case 8", args: args{A: 100, B: 77}, want: 188},
		{name: "Test Case 9", args: args{A: 1000, B: 777}, want: 1888},
		{name: "Test Case 10", args: args{A: 10000, B: 7777}, want: 18888},
		{name: "Test Case 11", args: args{A: 100000, B: 77777}, want: 188888},
		{name: "Test Case 12", args: args{A: 1000000, B: 777777}, want: 1888888},
		{name: "Test Case 13", args: args{A: 10000000, B: 7777777}, want: 18888888},
		{name: "Test Case 14", args: args{A: 100000000, B: 77777777}, want: 188888888},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeIt(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("MakeIt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMakeItLarge-8   	1000000000	         0.2569 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMakeItLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeIt(100000000, 77777777)
	}
}

// BenchmarkMakeItSmall-8   	1000000000	         0.2545 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMakeItSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeIt(1, 2)
	}
}

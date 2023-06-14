package numberofopendoors

import "testing"

func TestNumberOfOpenDoors(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5}, 2},
		{"test2", args{6}, 2},
		{"test3", args{1000}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfOpenDoors(tt.args.A); got != tt.want {
				t.Errorf("NumberOfOpenDoors() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 62883194                21.26 ns/op            0 B/op          0 allocs/op
func BenchmarkNumberOfOpenDoors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfOpenDoors(1000)
	}
}

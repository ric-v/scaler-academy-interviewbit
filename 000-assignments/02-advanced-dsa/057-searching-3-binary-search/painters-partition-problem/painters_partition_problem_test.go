package painterspartitionproblem

import "testing"

func TestPaintersPartitionProblem(t *testing.T) {
	type args struct {
		A int
		B int
		C []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2, 5, []int{1, 10}}, 50},
		{"test2", args{10, 1, []int{1, 8, 11, 3}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaintersPartitionProblem(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("PaintersPartitionProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPaintersPartitionProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PaintersPartitionProblem(10, 1, []int{1, 8, 11, 3})
	}
}

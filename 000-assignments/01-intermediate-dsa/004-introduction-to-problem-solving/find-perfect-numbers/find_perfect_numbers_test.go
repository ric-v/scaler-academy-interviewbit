package findperfectnumbers

import "testing"

func TestFindPerfectNumbers(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Test 1", args{6}, 6},
		{"Test 2", args{28}, 28},
		{"Test 3", args{496}, 496},
		{"Test 4", args{8128}, 8128},
		{"Test 5", args{33550336}, 33550336},
		{"Test 6", args{8589869056}, 8589869056},
		{"Test 7", args{137438691328}, 137438691328},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := FindPerfectNumbers(tt.args.A); gotResult != tt.wantResult {
				t.Errorf("FindPerfectNumbers() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkFindPerfectNumbers-8                  1        72427362108 ns/op              8 B/op          1 allocs/op
// BenchmarkFindPerfectNumbers-8               5713            192247 ns/op               0 B/op          0 allocs/op
func BenchmarkFindPerfectNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPerfectNumbers(137438691328)
	}
}

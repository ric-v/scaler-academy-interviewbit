package excelcolumntitle

import "testing"

func TestExcelColTitle(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{1}, "A"},
		{"test2", args{2}, "B"},
		{"test3", args{26}, "Z"},
		{"test4", args{27}, "AA"},
		{"test5", args{28}, "AB"},
		{"test6", args{52}, "AZ"},
		{"test7", args{53}, "BA"},
		{"test8", args{54}, "BB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExcelColTitle(tt.args.A); got != tt.want {
				t.Errorf("ExcelColTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkExcelColTitle-8   	 2021751	       566.9 ns/op	      32 B/op	       6 allocs/op
func BenchmarkExcelColTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExcelColTitle(1000000000)
	}
}

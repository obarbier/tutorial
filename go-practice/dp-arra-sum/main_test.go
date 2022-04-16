package main

import "testing"

func Test_maxSubsetSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"test1",
			args{
				arr: []int32{-2, 1, 3, -4, 5},
			},
			8},
		{"test2",
			args{
				arr: []int32{2, 1, 5, 8, 4},
			},
			11},
		{
			name: "test3",
			args: args{arr: []int32{3, 7, 4, 6, 5}},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsetSum(tt.args.arr); got != tt.want {
				t.Errorf("maxSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

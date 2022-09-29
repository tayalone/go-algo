package dynamicprograming

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	type args struct {
		// mat [][]int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		// want [][]int
		want [][]int
	}{
		{
			name: "case 3",
			args: args{
				mat: [][]int{{0}, {1}, {0}},
			},
			want: [][]int{{0}, {1}, {0}},
		},
		{
			name: "case 4",
			args: args{
				mat: [][]int{{0}, {1}, {1}, {0}},
			},
			want: [][]int{{0}, {1}, {1}, {0}},
		},
		{
			name: "case 5",
			args: args{
				mat: [][]int{{0}, {1}, {1}, {1}, {0}},
			},
			want: [][]int{{0}, {1}, {2}, {1}, {0}},
		},
		// // ----------  horizontal ------ // //
		// // TODO: Add test cases.
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		mat: []int{0},
		// 	},
		// 	want: []int{0},
		// },
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		mat: []int{0, 0},
		// 	},
		// 	want: []int{0, 0},
		// },
		// {
		// 	name: "case 3",
		// 	args: args{
		// 		mat: []int{0, 1, 0},
		// 	},
		// 	want: []int{0, 1, 0},
		// },
		// {
		// 	name: "case 4",
		// 	args: args{
		// 		mat: []int{0, 1, 1, 0},
		// 	},
		// 	want: []int{0, 1, 1, 0},
		// },
		// {
		// 	name: "case 5",
		// 	args: args{
		// 		mat: []int{0, 1, 1, 1, 0},
		// 	},
		// 	want: []int{0, 1, 2, 1, 0},
		// },
		// {
		// 	name: "case 6",
		// 	args: args{
		// 		mat: []int{0, 1, 1, 1, 1, 0},
		// 	},
		// 	want: []int{0, 1, 2, 2, 1, 0},
		// },
		// {
		// 	name: "case 7",
		// 	args: args{
		// 		mat: []int{0, 1, 1, 1, 1, 1, 0},
		// 	},
		// 	want: []int{0, 1, 2, 3, 2, 1, 0},
		// },
		// // -------------------------------------
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

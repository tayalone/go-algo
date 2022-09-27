package dfs

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		// TODO: Add test cases.
		{
			name: "Case 0",
			args: args{
				sr:    1,
				sc:    0,
				color: 2,
				image: [][]int{{0, 0, 0}, {0, 0, 0}},
			},
			want: [][]int{{2, 2, 2}, {2, 2, 2}},
		},
		{
			name: "Case 1",
			args: args{
				sr:    1,
				sc:    1,
				color: 2,
				image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			},
			want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			name: "Case 2",
			args: args{
				sr:    0,
				sc:    0,
				color: 0,
				image: [][]int{{0, 0, 0}, {0, 0, 0}},
			},
			want: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trans(t *testing.T) {
	type args struct {
		image    *[][]int
		sr       int
		sc       int
		ogValue  int
		newValue int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trans(tt.args.image, tt.args.sr, tt.args.sc, tt.args.ogValue, tt.args.newValue)
		})
	}
}

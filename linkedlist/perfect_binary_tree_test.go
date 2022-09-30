package linkedlist

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/tayalone/go-algo/ds"
)

func bs(root *ds.PerfectNode) []string {
	if root == nil {
		return []string{"#"}
	}
	c := []string{strconv.Itoa(root.Val)}
	n := bs(root.Next)
	return append(c, n...)
}

func ldfs(root *ds.PerfectNode) []string {
	b := bs(root)
	if root.Left != nil {
		l := ldfs(root.Left)
		return append(b, l...)
	}
	return b
}

func makeWantPBT(root *ds.PerfectNode) []string {
	if root == nil {
		return []string{}
	}
	return ldfs(root)
}

func TestConnect(t *testing.T) {
	type args struct {
		root *ds.PerfectNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case 0",
			args: args{
				root: &ds.PerfectNode{
					Val:  -1,
					Next: nil,
					Left: &ds.PerfectNode{
						Val:  0,
						Next: nil,
						Left: &ds.PerfectNode{
							Val:  2,
							Next: nil,
							Left: &ds.PerfectNode{
								Val:   6,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
							Right: &ds.PerfectNode{
								Val:   7,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
						},
						Right: &ds.PerfectNode{
							Val:  3,
							Next: nil,
							Left: &ds.PerfectNode{
								Val:   8,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
							Right: &ds.PerfectNode{
								Val:   9,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &ds.PerfectNode{
						Val:  1,
						Next: nil,
						Left: &ds.PerfectNode{
							Val:  4,
							Next: nil,
							Left: &ds.PerfectNode{
								Val:   10,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
							Right: &ds.PerfectNode{
								Val:   11,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
						},
						Right: &ds.PerfectNode{
							Val:  5,
							Next: nil,
							Left: &ds.PerfectNode{
								Val:   12,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
							Right: &ds.PerfectNode{
								Val:   13,
								Next:  nil,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
			want: []string{
				"-1", "#", "0", "1", "#", "2", "3", "4", "5", "#", "6", "7", "8", "9", "10", "11", "12", "13", "#",
			},
		},
		{
			name: "case 1",
			args: args{
				root: &ds.PerfectNode{
					Val:  1,
					Next: nil,
					Left: &ds.PerfectNode{
						Val:  2,
						Next: nil,
						Left: &ds.PerfectNode{
							Val:   4,
							Next:  nil,
							Left:  nil,
							Right: nil,
						},
						Right: &ds.PerfectNode{
							Val:   5,
							Next:  nil,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &ds.PerfectNode{
						Val:  3,
						Next: nil,
						Left: &ds.PerfectNode{
							Val:   6,
							Next:  nil,
							Left:  nil,
							Right: nil,
						},
						Right: &ds.PerfectNode{
							Val:   7,
							Next:  nil,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: []string{"1", "#", "2", "3", "#", "4", "5", "6", "7", "#"},
		},
		{
			name: "case 2",
			args: args{
				root: nil,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nt := Connect(tt.args.root)
			got := makeWantPBT(nt)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

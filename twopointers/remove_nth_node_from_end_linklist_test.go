package twopointers

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "CASE_11",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 1,
			},
			want: []int{2, 3},
		},
		{
			name: "CASE_10",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 2,
			},
			want: []int{1, 3},
		},
		{
			name: "CASE_9",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 3,
			},
			want: []int{2, 3},
		},
		{
			name: "CASE_8",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 1,
			},
			want: []int{1},
		},
		{
			name: "CASE_7",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 2,
			},
			want: []int{2},
		},
		{
			name: "CASE_6",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				n: 4,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "CASE_5",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				n: 3,
			},
			want: []int{1, 3, 4},
		},
		{
			name: "CASE_4",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				n: 2,
			},
			want: []int{1, 2, 4},
		},
		{
			name: "CASE_3",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				n: 1,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "CASE_2",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 0,
			},
			want: []int{},
		},
		{
			name: "CASE_1",
			args: args{
				head: nil,
				n:    0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpGot := RemoveNthFromEnd(tt.args.head, tt.args.n)
			got := []int{}
			for tmpGot != nil {
				got = append(got, tmpGot.Val)
				tmpGot = tmpGot.Next
			}

			if !reflect.DeepEqual([]int{}, []int{}) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

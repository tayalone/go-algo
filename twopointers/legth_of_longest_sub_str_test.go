package twopointers

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Case 1",
		// 	args: args{
		// 		s: "abcabcbb",
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Case 2",
		// 	args: args{
		// 		s: "bbbbb",
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "Case 3",
		// 	args: args{
		// 		s: "pwwkew",
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Case 4",
		// 	args: args{
		// 		s: "a",
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "Case 5",
		// 	args: args{
		// 		s: "aa",
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "Case 5",
		// 	args: args{
		// 		s: "ab",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "Case 6",
		// 	args: args{
		// 		s: "dvdf",
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "Case 6",
		// 	args: args{
		// 		s: "anviaj",
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "Case 7",
		// 	args: args{
		// 		s: "tmmzuxt",
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "Case 8",
		// 	args: args{
		// 		s: "ohvhjdml",
		// 	},
		// 	want: 6,
		// },
		{
			name: "Case 9",
			args: args{
				s: "qrsvbspk",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

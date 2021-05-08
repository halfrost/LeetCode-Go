package leetcode

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	cases := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				commands:  []int{4, -1, 3},
				obstacles: [][]int{{}},
			},
			25,
		},
		{
			"case 2",
			args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			65,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}

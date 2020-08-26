package structures

import (
	"reflect"
	"testing"
)

func Test_Intss2Points(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			"测试 [][]int 转换成 []Point ",
			args{
				[][]int{
					{1, 0},
					{2, 0},
					{3, 0},
					{4, 0},
					{5, 0},
				},
			},
			[]Point{
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		if got := Intss2Points(tt.args.points); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. intss2Points() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_Points2Intss(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"测试 [][]int 转换成 []Point ",
			args{
				[]Point{
					{X: 1, Y: 0},
					{X: 2, Y: 0},
					{X: 3, Y: 0},
					{X: 4, Y: 0},
					{X: 5, Y: 0},
				},
			},
			[][]int{
				{1, 0},
				{2, 0},
				{3, 0},
				{4, 0},
				{5, 0},
			},
		},
	}
	for _, tt := range tests {
		if got := Points2Intss(tt.args.points); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Points2Intss() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

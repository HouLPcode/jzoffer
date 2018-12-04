package test

import (
	"github.com/HouLPcode/jzoffer/alg"
	"testing"
)

type Data13Type struct {
	Threshold int
	Rows      int
	Cols      int
	Expected  int
}

var Test13Datas = []Data13Type{
	{5, 10, 10, 21},
	{15, 20, 20, 359},
	{10, 1, 100, 29},
	{10, 1, 10, 10},
	{15, 100, 1, 79},
	{15, 10, 1, 10},
	{15, 1, 1, 1},
	{0, 1, 1, 1},
	{-10, 10, 10, 0},
}

func Test13(t *testing.T) {
	run13Tests(t, func(threshold, rows, cols int) int {
		cnt := alg.MovingCount_13(threshold, rows, cols)
		return cnt
	})
}

func run13Tests(t *testing.T, f func(threshold, rows, cols int) int) {
	for i, test := range Test13Datas {
		output := f(test.Threshold, test.Rows, test.Cols)
		if output != test.Expected {
			t.Errorf("test %d: output mismatch:\ngot %d\nwant %d\n", i, output, test.Expected)
		}
	}
}

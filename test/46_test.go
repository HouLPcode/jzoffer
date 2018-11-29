package test

import (
	"github.com/HouLPcode/jzoffer/alg"
	"testing"
)

type dataTest struct {
	val    int
	output int
	error  string
}

var dataTests = []dataTest{
	{val: 0, output: 1},
	{val: 10, output: 2},
	{val: 125, output: 3},
	{val: 126, output: 2},
	{val: 426, output: 1},
	{val: 100, output: 2},
	{val: 101, output: 2},
	{val: 12258, output: 5},
	{val: -100, output: -1},
}

func Test46(t *testing.T) {
	run46Tests(t, func(val int) (int, error) {
		n := alg.GetTranslationCount(val)
		return n, nil
	})
}

func run46Tests(t *testing.T, f func(val int) (int, error)) {
	for i, test := range dataTests {
		output, err := f(test.val)
		if err == nil && output != test.output {
			t.Errorf("test %d: output mismatch:\ngot   %d\nwant  %d\ninput %d\n",
				i, output, test.output, test.val)
		}
	}
}

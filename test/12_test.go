package test

import (
	"errors"
	"github.com/HouLPcode/jzoffer/alg"
	"testing"
)

type Data12Type struct {
	Matrix string
	Rows   int
	Cols   int
	Path   string
	Expect bool
	Err    error
}

var Test12Datas = []Data12Type{
	{"ABTGCFCSJDEH", 3, 4, "BFCE", true, errors.New("")},
	{"ABCESFCSADEE", 3, 4, "SEE", true, errors.New("")},
	{"ABTGCFCSJDEH", 3, 4, "ABFB", false, errors.New("")},
	{"ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SLHECCEIDEJFGGFIE", true, errors.New("")},
	{"ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEHJIGQEM", true, errors.New("")},
	{"ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEEJIGOEM", false, errors.New("")},
	{"ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEHJIGQEMS", false, errors.New("")},
	{"AAAAAAAAAAAA", 3, 4, "AAAAAAAAAAAA", true, errors.New("")},
	{"AAAAAAAAAAAA", 3, 4, "AAAAAAAAAAAAA", false, errors.New("")},
	{"A", 1, 1, "A", true, errors.New("")},
	{"A", 1, 1, "B", false, errors.New("")},
	{"", 0, 0, "", false, errors.New("nil input err")},
}

func Test12(t *testing.T) {
	run12Tests(t, func(matrix string, rows, cols int, str string) bool {
		b := alg.HasPath(matrix, rows, cols, str)
		return b
	})
}

func run12Tests(t *testing.T, f func(matrix string, rows, cols int, str string) bool) {
	for i, test := range Test12Datas {
		output := f(test.Matrix, test.Rows, test.Cols, test.Path)
		if output != test.Expect {
			// TODO 怎么输出bool类型
			t.Errorf("test %d: output mismatch:\ngot %t\nwant %t\n", i, output, test.Expect)
		}
	}
}

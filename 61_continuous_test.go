package jzoffer

import (
	"testing"
)

// 面试题61：扑克牌的顺子
// 题目：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
// 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字。

// 桶排序，0表示大小王
func isContinuous(data []int) bool {
	buf := make([]bool, 14)
	cnt0 := 0
	for _, v := range data {
		if v == 0 {
			cnt0++
		} else {
			if buf[v] == true {
				return false //有重牌
			}
			buf[v] = true
		}
	}
	// 找最小值
	i := 1
	for ; i < 14 && buf[i] == false; i++ {
	}
	for j := i; j < i+5 && j < 14; j++ {
		if buf[j] == false {
			cnt0--
		}
		if cnt0 < 0 {
			return false
		}
	}
	return true
}

func Test_IsContinuous(t *testing.T) {
	type testType struct {
		d []int
		f bool
	}
	testCase := []testType{
		testType{[]int{1, 3, 2, 5, 4}, true},
		testType{[]int{1, 3, 2, 6, 4}, false},
		testType{[]int{0, 3, 2, 6, 4}, true},
		testType{[]int{0, 3, 1, 6, 4}, false},
		testType{[]int{1, 3, 0, 5, 0}, true},
		testType{[]int{1, 3, 0, 7, 0}, false},
		testType{[]int{1, 0, 0, 5, 0}, true},
		testType{[]int{1, 0, 0, 7, 0}, false},
		testType{[]int{3, 0, 0, 0, 0}, true},
		testType{[]int{0, 0, 0, 0, 0}, true},
		testType{[]int{1, 0, 0, 1, 0}, false},
	}

	for _, test := range testCase {
		ans := isContinuous(test.d)
		if test.f != ans {
			t.Error("测试失败，测试集", test, "实际输出", ans)
		}
	}
}

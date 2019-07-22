package jzoffer

import (
	"testing"
)

// 面试题62：圆圈中最后剩下的数字
// 题目：0, 1, …, n-1这n个数字排成一个圆圈，从数字0开始每次从这个圆圈里
// 删除第m个数字。求出这个圆圈里剩下的最后一个数字。

// f(n) = 0, n=1
// f(n) = (f(n-1) + m) % n, n>1
func LastRemaining(n, m int) int {
	if n <= 0 || m <= 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	rtn := 0
	for i := 2; i <= n; i++ { // 此处必须从小到大遍历
		rtn = (rtn + m) % i
	}
	return rtn
}

func Test_LastRemaining(t *testing.T) {
	testCase := [][]int{
		[]int{5, 3, 3},
		[]int{5, 2, 2},
		[]int{6, 7, 4},
		[]int{6, 6, 3},
		[]int{0, 0, -1},
		[]int{4000, 997, 1027}}

	for _, test := range testCase {
		ans := LastRemaining(test[0], test[1])
		if test[2] != ans {
			t.Error("测试失败，测试集", test, "实际输出", ans)
		}
	}
}

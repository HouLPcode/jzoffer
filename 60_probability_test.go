package jzoffer

import (
	"fmt"
	"testing"
)

// 面试题60：n个骰子的点数
// 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s
// 的所有可能的值出现的概率。

// slice统计计数
// 错误思路：n控制的是循环的层数，所以不要一层一层实现，遍历每种可能
// 解法，f(n,s) = f(n-1,s-1)+f(n-1,s-2).....+f(n-1,s-6)
// 即和为s的种类数是n-1的筛子前6中可能
// 两个数组交替使用
func probability(n int) {
	if n < 1 {
		return
	}
	// 总个数
	// total := int(math.Pow(6, n))
	buf := make([][]int, 2)
	buf[0] = make([]int, 6*n+1) // 下标对应和，内容对应次数
	buf[1] = make([]int, 6*n+1)

	flag := 0
	for i := 1; i <= 6; i++ { // 1个骰子的情况
		buf[flag][i] = 1 // 1个骰子和为1-6，次数均是1次
	}
	for k := 2; k <= n; k++ { // 2到n个骰子
		// k个骰子的和为 k~6k
		for i := k; i <= 6*k; i++ {
			buf[1-flag][i] = 0
			// 前n-1,n-2,n-3.....n-6的可能
			for j := 1; j <= 6 && i-j >= 1; j++ { // 第k个筛子的值为1-6，和为i的值的次数是前6项的次数和，最多6次
				buf[1-flag][i] += buf[flag][i-j]
			}
		}
		flag = 1 - flag // 数组交替
	}
	// 最终取范围 n...6n的值
	for i := n; i <= 6*n; i++ {
		fmt.Println(i, buf[flag][i])
	}
}

func Test_probability(t *testing.T) {
	fmt.Println("1个骰子")
	probability(1)
	fmt.Println("2个骰子")
	probability(2)
	fmt.Println("3个骰子")
	probability(3)
}

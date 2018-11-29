// 面试题46：把数字翻译成字符串
// 题目：给定一个数字，我们按照如下规则把它翻译为字符串：0翻译成"a"，1翻
// 译成"b"，……，11翻译成"l"，……，25翻译成"z"。一个数字可能有多个翻译。例
// 如12258有5种不同的翻译，它们分别是"bccfi"、"bwfi"、"bczi"、"mcfi"和
// "mzi"。请编程实现一个函数用来计算一个数字有多少种不同的翻译方法。

//递归，从右向左，避免子问题中的重复
package alg

import (
	"strconv"
)

// TODO 负数返回值需要自行商定 -100 -> 0

func GetTranslationCount(num int) int {
	if num < 0 {
		return -1
	}
	s := strconv.Itoa(num)
	t := GetTranslationCountStr(s)
	return t
}

// 对于 12258：
// f(4) = 1
// f(3) = f(4)+0 = 1
// f(2) = f(3)+f(4) = 2
// f(1) = f(2)+f(3) = 3
// f(0) = f(1)+f(2) = 5
func GetTranslationCountStr(s string) int {
	l := len(s)
	if l == 1 {
		return 1
	}
	tmp := make([]int, l, l)
	tmp[l-1] = 1
	tmp[l-2] = tmp[l-1] + tflag(s[l-2:])
	for i := l - 3; i >= 0; i-- {
		tmp[i] = tmp[i+1] + tmp[i+2]*tflag(s[i:i+2])
	}
	return tmp[0]
}

func tflag(s string) int {
	num, _ := strconv.Atoi(s)
	if num > 9 && num < 26 {
		return 1
	}
	return 0
}

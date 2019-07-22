package jzoffer

import (
	"testing"
)

// 50_01_FirstNotRepeatingChar
// 面试题50（一）：字符串中第一个只出现一次的字符
// 题目：在字符串中找出第一个只出现一次的字符。如输入"abaccdeff"，则输出
// 'b'。
// 类似桶排序
func firstNotRepeatingChar(pString string) byte {
	bucket := make([]int, 128)
	for i := 0; i < len(pString); i++ {
		bucket[int(pString[i])]++
	}
	i := 0
	for ; i < len(pString); i++ {
		if bucket[int(pString[i])] == 1 {
			break
		}
	}
	return pString[i]
}

var testCase = []struct {
	in  string
	out byte
}{
	{
		"google", // 常规输入测试，存在只出现一次的字符
		'l',
	},
	// {
	// 	"aabccdbd", // 常规输入测试，不存在只出现一次的字符
	// 	'\0'
	// },
	{
		"abcdefg", // 常规输入测试，所有字符都只出现一次
		'a',
	},
	// {
	// 	"", // 鲁棒性测试，输入nullptr
	// 	'\0'
	// },
}

func Test_firstNotRepeatingChar(t *testing.T) {
	for _, test := range testCase {
		ans := firstNotRepeatingChar(test.in)
		if test.out != ans {
			t.Error("测试失败，测试集", test, "实际输出", ans)
		}
	}
}

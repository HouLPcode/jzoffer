package alg

// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

//path:BFCE
func HasPath(matrix string, rows, cols int, path string) bool {
	//strlen := len(path)

	plen := 0 //查找str字符串的第多少位，前面的字符串已经全部找到
	visited := make([]bool, rows*cols, rows*cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			//TODO 数组参数传递
			b := hasPathCore(matrix, rows, cols, row, col, &plen, path, &visited)
			if b == true {
				return b
			}
		}
	}
	return false
}

func hasPathCore(matrix string, rows, cols, row, col int, plen *int, path string, visited *[]bool) bool {
	if *plen == len(path) { //?????????????
		return true
	}

	if row >= 0 && row < rows && col >= 0 && col < cols && //范围合法
		(*visited)[row*cols+col] == false && // 没有访问过
		path[*plen] == matrix[row*cols+col] { // 字符匹配
		*plen++
		(*visited)[row*cols+col] = true
		var has bool
		has = hasPathCore(matrix, rows, cols, row+1, col, plen, path, visited) ||
			hasPathCore(matrix, rows, cols, row, col+1, plen, path, visited) ||
			hasPathCore(matrix, rows, cols, row-1, col, plen, path, visited) ||
			hasPathCore(matrix, rows, cols, row, col-1, plen, path, visited)
		if has == false {
			*plen--
			(*visited)[row*cols+col] = false
		}
		return has
	}

	return false
}

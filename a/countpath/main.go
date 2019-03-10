package main

import "fmt"

func main() {
	// 0 表示空地
	// 1 表示石头
	grid := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 1, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{1, 0, 1, 0, 0, 1, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 1, 0, 1, 0},
		[]int{0, 1, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println("countPaths(grid) =", countPaths(grid))
}

func countPaths(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	// 建一个二维数组
	opt := make([][]int, row)
	for k := range opt {
		opt[k] = make([]int, col)
	}

	// 最下面一行，只有一种走法。
	for idx := 0; idx < col; idx++ {
		opt[row-1][idx] = 1
	}

	// 最右边那一列，也只有一种走法。
	for idx := 0; idx < row; idx++ {
		opt[idx][col-1] = 1
	}

	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			if grid[i][j] == 0 {
				// 是空地
				opt[i][j] = opt[i+1][j] + opt[i][j+1]

			} else {
				// 是石头
				opt[i][j] = 0
			}
		}
	}

	return opt[0][0]
}

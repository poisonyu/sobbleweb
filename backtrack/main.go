package main

import (
	"fmt"
	"sort"
)

// 子集和问题

// 给定一个正整数数组nums和一个目标正整数target，找出所有可能的组合，使得组合的元素和等于target
// 给定数组无重复元素，每个元素可以被选取多次，以列表的形式返回组合，列表中不包含重复组合

//	func backtrackSubsetSumI(total, target int, state *[]int, nums *[]int, res *[][]int) {
//		if total == target {
//			*res = append(*res, append([]int{}, *state...))
//			return
//		}
//		for _, num := range *nums {
//			if total+num > target || len(*state) > 0 && (*state)[len(*state)-1] > num {
//				continue
//			}
//			*state = append(*state, num)
//			backtrackSubsetSumI(total+num, target, state, nums, res)
//			*state = (*state)[:len(*state)-1]
//		}
//	}
func backtrackSubsetSumI(start, target int, state *[]int, nums *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *state...))
		return
	}
	for i := start; i < len(*nums); i++ {
		num := (*nums)[i]
		if target-num < 0 {
			break
		}
		*state = append(*state, num)
		// 每个元素可以被多次选取 start = i
		backtrackSubsetSumI(i, target-num, state, nums, res)
		*state = (*state)[:len(*state)-1]
	}
}

// 给定数组可能包含重复元素，每个元素只可被选择一次，以列表的形式返回组合，列表中不包含重复组合
func backtrackSubsetSumII(start, target int, state *[]int, nums *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *state...))
		return
	}
	for i := start; i < len(*nums); i++ {
		num := (*nums)[i]
		if target-num < 0 {
			break
		}
		// 去重列表中的重复组合
		// 由于nums已经被排序，相邻元素可能相等
		if i > start && (*nums)[i-1] == num {
			continue
		}
		*state = append(*state, num)
		// 每个元素只能被选择一次 start = i+1
		backtrackSubsetSumI(i+1, target-num, state, nums, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSumI(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	var start int
	sort.Ints(nums)
	backtrackSubsetSumI(start, target, &state, &nums, &res)
	return res
}

func subsetSumII(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	var start int
	sort.Ints(nums)
	backtrackSubsetSumII(start, target, &state, &nums, &res)
	return res
}

// N皇后问题

func backtrack(row, n int, cols, diags1, diags2 *[]bool, state *[][]string, res *[][][]string) {
	if row == n {

		newState := make([][]string, n)
		for i, _ := range *state {
			row := make([]string, n)
			copy(row, (*state)[i])
			newState[i] = row
		}
		*res = append(*res, newState)
		return
	}

	for col := 0; col < n; col++ {
		diag1 := row - col + n - 1
		diag2 := row + col
		if !(*cols)[col] && !(*diags1)[diag1] && !(*diags2)[diag2] {
			(*state)[row][col] = "Q"
			(*cols)[col], (*diags1)[diag1], (*diags2)[diag2] = true, true, true
			backtrack(row+1, n, cols, diags1, diags2, state, res)
			(*state)[row][col] = "#"
			(*cols)[col], (*diags1)[diag1], (*diags2)[diag2] = false, false, false
		}
	}
}

func nQueens(n int) [][][]string {
	// 初始化棋盘
	state := make([][]string, n)
	// #表示空，Q表示queen
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "#"
		}
		state[i] = row
	}
	cols := make([]bool, n)
	diags1 := make([]bool, 2*n-1)
	diags2 := make([]bool, 2*n-1)
	res := make([][][]string, 0)
	backtrack(0, n, &cols, &diags1, &diags2, &state, &res)
	return res
}
func main() {
	// a := []int{4, 5, 5}
	// fmt.Println(subsetSumII(a, 10))
	fmt.Println(nQueens(4))

}

package main

import "fmt"

// 给定一个共有n阶的楼梯，每步可以上1阶或者2阶，请问有多少种方案可以爬到楼顶

func backtrack(target int, state *[]int, nums *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *state...))
		return
	}

	for _, num := range *nums {
		if target-num < 0 {
			// nums排序，用break,没排序，用continue
			continue
		}
		*state = append(*state, num)
		backtrack(target-num, state, nums, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSum(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	backtrack(target, &state, &nums, &res)
	return res
}

// 暴力搜索
func dfs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return dfs(n-1) + dfs(n-2)
}

func climbingStairsDFS(n int) int {
	return dfs(n)
}

// 记忆化搜索
// 重叠子问题都只被计算一次

func dfsMem(i int, mem []int) int {
	if i == 1 || i == 2 {
		return i
	}
	fmt.Printf("dp[%d]: %v\n", i, mem)
	// dp[i]存在，直接返回
	if mem[i] != -1 {
		return mem[i]
	}
	// dp[i]不存在，递归计算dp[i]
	count := dfsMem(i-1, mem) + dfsMem(i-2, mem)
	// 记录dp[i]
	mem[i] = count
	return count
}

func climbingStairsDFSMem(n int) int {
	mem := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		mem[i] = -1
	}
	return dfsMem(n, mem)
}

// 动态规划

func main() {
	// a := []int{1, 2}
	// fmt.Println(subsetSum(a, 3))
	// fmt.Println(climbingStairsDFS(5))
	fmt.Println(climbingStairsDFSMem(5))
}

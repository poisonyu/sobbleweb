package main

import "fmt"

func main() {
	a := []int{3, 4, 1, 2, 10, 5, 11, 6}
	fmt.Println(a)
	selectSrot(a)
	fmt.Println(a)
}

// 选择排序
func selectSrot(s []int) []int {
	first := 0
	for first < len(s)-1 {
		k := first // k当前未排序区间最小元素的索引
		// j循环未排序区间,寻找区间中最小元素，并把元素赋值给k
		for j := first + 1; j < len(s); j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		s[first], s[k] = s[k], s[first]
		first++
	}
	return s
}

// 时间复杂度 平方阶
// 非自适应排序 不能利用数据已有的顺序信息减少计算量
// 空间复杂度 常数阶 原地排序 指针i, j使用常数大小的额外空间
// 非稳定排序 相等元素的相对顺序在排序后会发生变化

/* 选择排序 */
func selectionSort(nums []int) {
	n := len(nums)
	// 外循环：未排序区间为 [i, n-1]
	for i := 0; i < n-1; i++ {
		// 内循环：找到未排序区间内的最小元素
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				// 记录最小元素的索引
				k = j
			}
		}
		// 将该最小元素与未排序区间的首个元素交换
		nums[i], nums[k] = nums[k], nums[i]

	}
}

// func selectSrot(s []int) []int {
// 	first := 0
// 	for first < len(s)-1 {
// 		minIndex, _ := min(s[first:])
// 		s[first], s[minIndex+first] = s[minIndex+first], s[first]
// 		first++
// 	}
// 	return s
// }

// func min(s []int) (minIndex, min int) {
// 	minIndex = 0
// 	min = s[0]
// 	for i := 1; i < len(s); i++ {
// 		if s[i] < min {
// 			min = s[i]
// 			minIndex = i
// 		}
// 	}
// 	return
// }

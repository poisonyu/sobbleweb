package main

// 冒泡排序
// 连续比较交换相邻元素实现排序

// 时间复杂度 平方阶 各轮冒泡遍历数组的长度依次为n-1、n-2、...、2、1,总和为(n-1)n/2
// 自适应排序 能够用数组已有的顺序信息减少计算量
// 空间复杂度 常数阶 原地排序 直接在原数组上操作实现排序
// 稳定排序 相等元素在排序后的相对位置不变

func bubbleSort(nums []int) {
	n := len(nums)
	swap := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

package main

import (
	"fmt"
)

// 归并排序

func mergeSort(nums []int, left, right int) {
	// 子切片长度不大于1时，终止
	if left >= right {
		return
	}
	// 当前切片的中点索引
	mid := left + (right-left)/2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	// 双指针和临时切片的索引
	i, j, k := left, mid+1, 0
	// 创建一个临时的切片
	// nums是整个要排序的切片，当前合并的切片长度由left, right决定
	tmp := make([]int, right-left+1)
	// for i < j {
	// 	for i <= mid && nums[i] < nums[j] {
	// 		tmp[k] = nums[i]
	// 		i++
	// 		k++
	// 	}
	// 	for j <= right && nums[j] < nums[i] {
	// 		tmp[k] = nums[j]
	// 		j++
	// 		k++
	// 	}
	// }

	//
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	//

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	//
	// nums是整个要排序的切片，当前合并的切片长度由left, right决定
	// 所以将切片复制到nums时需要加上left定位
	for k := 0; k < len(tmp); k++ {
		nums[left+k] = tmp[k]
	}
	//
	// for i := 0; i < len(nums); i++ {
	// 	nums[i] = tmp[i]
	// }
}

func main() {
	a := []int{3, 4, 1, 2, 10, 5, 11, 6, 12, 15, 22, 6, 8}
	fmt.Println(a)
	mergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}

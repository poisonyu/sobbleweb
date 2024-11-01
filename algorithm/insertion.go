package main

import (
	"fmt"
)

// 插入排序
func insertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		base := s[i]
		// 内层循环是倒着来的
		// 从已排序的结果中选最大值与要排序的值base比较，如果比base大，
		// 当前最大值向后移动一位
		// 然后就选下一个已排序的值与base比较
		j := i - 1
		for j >= 0 && s[j] > base {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = base
	}
}

// 尾递归
func (q *quickSortTailCall) quickSort(nums int, left, right int) {
	for left < right {
		pivot := q.partition(nums, left, right)
		if (pivot - left) < (right - pivot) {
			q.quickSort(nums, left, pivot-1)
			left = pivot + 1
		} else {
			q.quickSort(nums, pivot+1, right)
			right = pivot - 1
		}
	}
}

func main() {
	a := []int{3, 4, 1, 2, 10, 5, 11, 6}
	fmt.Println(a)

	qs := new(quickSortMedian)
	qs.quickSort(a, 0, len(a)-1)

	// insertionSort(a)
	fmt.Println(a)
}

// func insertion(s []int) {
// 	for i := 1; i < len(s); i++ {
// 		a := s[i]
// 		for j := 0; j < i; j++ {
// 			if s[j] > a {
// 				s[j+1 : i+1] = s[j:i]
// 				s[j] = a
// 				break
// 			}
// 		}
// 	}
// }

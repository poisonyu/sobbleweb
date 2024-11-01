package main

// 快速排序
type quickSort struct{}

// 哨兵划分
func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

// 分治递归
func (q *quickSort) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

// 基准数优化
type quickSortMedian struct{}

func (q *quickSortMedian) partition(nums []int, left, right int) int {
	med := q.medianThree(nums, left, (left+right)/2, right)
	nums[left], nums[med] = nums[med], nums[left]
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

// 获取一个合适的哨兵划分的索引
func (q *quickSortMedian) medianThree(nums []int, left, mid, right int) int {
	l, m, r := nums[left], nums[mid], nums[right]
	// m在中间
	if (m > r && m < l) || (m > l && m < r) {
		return mid
		// r在中间
	} else if (r > m && r < l) || (r > l && r < m) {
		return right
	} else {
		return left
	}
}

func (q *quickSortMedian) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

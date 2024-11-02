package main

// 快速排序
// 基于分治策略的排序算法
// 快速排序的核心操作时哨兵划分，选择数组中的一个元素为基准数，
// 将所有小于基准数的元素移到其左侧，大于基准数的元素移到其右侧
// 哨兵划分的实质是将一个较长数组的排序问题简化为两个较短数组的排序问题

// 时间复杂度 线性对数阶 哨兵划分的递归层数为logn
// (在最差的情况下，哨兵划分产生长度为0和n-1的两个子数组，此时递归层数达到n),
// 每层中的总循环数为n，总体使用时间为nlogn
// 非自适应性排序
// 空间复杂度 线性阶 原地排序
// 非稳定排序 基准数会发生交换

type quickSort struct{}
type quickSortMedian struct{}
type quickSortTailCall struct{}

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

// 快速排序为什么快 与归并排序，堆排序的平均时间复杂度相同
// 出现最差情况的概率很低 快速排序在绝大多数情况下都是以线性对数阶的时间复杂度运行的
// 缓存使用效率高 在执行哨兵划分时，系统可以将整个子数组加载到缓存，访问元素的效率较高(堆排序需要跳跃式访问元素，缺乏这一特性)
// 复杂度的常数系数小 快速排序的比较、赋值、交换等操作的总数量最少

// 基准数优化
// 优化哨兵划分的基准数选取策略
// 获取一个合适的哨兵划分的基准数索引
// 选取首、尾、中三个元素的中位数为基准数
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

func (q *quickSortMedian) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

// 尾递归
// 在哨兵划分后，仅对较短的子数组进行递归
func (q *quickSortTailCall) quickSort(nums []int, left, right int) {
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

func (q *quickSortTailCall) partition(nums []int, left, right int) int {
	// todo
	return 1
}

# Hello Algorithm

## 11.6 归并排序(merge sort)
**基于分治策略**

分为**划分**和**合并**两个阶段

划分阶段通过递归不断将数组从中点处分开，将长数组排序问题转换为短数组排序问题

合并阶段 当子数组长度为1时终止划分，开始合并，将两个子数组有序合并成一个较长的有序数组，直至结束。

先递归左子数组，再递归右子数组，最后处理合并

时间复杂度 **线性对数阶** 划分产生高度为logn的递归树，每层合并的总操作量为n,总体时间复杂度为nlogn

空间复杂度 **线性阶** 

**非自适应排序**  
**非原地排序**  
**稳定排序**  

```
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
	// copy(nums, tmp)
}
```

## 11.7 堆排序(heap sort)
基于堆数据结构实现的高效排序算法  
**建堆操作**和**元素出堆操作**实现堆排序  
~~建立一个小顶堆，然后不断出堆元素，把元素存在额外的数组中~~  
建立一个大顶堆，将堆顶元素与堆底元素交换，交换后，堆的长度减1，已排序元素加1，从顶到底堆化，恢复堆的性质。

```
// 堆排序
func siftDown(nums *[]int, n, i int) {
	for {
		l, r, max := 2*i+1, 2*i+2, i
		for l < n && (*nums)[l] > (*nums)[max] {
			max = l
		}
		for r < n && (*nums)[r] > (*nums)[max] {
			max = r
		}
		if max == i {
			break
		}
		(*nums)[i], (*nums)[max] = (*nums)[max], (*nums)[i]
		i = max
	}
}

func heapSort(nums *[]int) {
	// 通过从顶至底堆化，将切片变成大顶堆
	// len(*nums)/2-1获取完全二叉树非叶节点的最小节点的索引
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(*nums), i)
	}
	for i := len(*nums) - 1; i > 0; i-- {
		// 交换最大元素
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		siftDown(nums, i, 0)
	}
}
```

时间复杂度 **线性对数阶** 建堆操作使用O(n)时间，从堆中提取最大元素的时间复杂度为O(logn),共循环n-1轮  
非自适应排序  
空间复杂度 **常数阶**   
原地排序  
非稳定排序

## 桶排序(bucket sort)
基于分治策略  非比较排序算法  
设置一些有大小顺序的桶，每个桶对应一个数据范围，将数据平均分配到各个桶中，分别对每个桶排序，再按桶的顺序合并所有数据。  

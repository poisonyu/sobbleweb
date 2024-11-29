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
```
func bucketSort(nums []float64) {
	// 初始化桶
	k := len(nums) / 2
	buckets := make([][]float64, k)
	for i := 0; i < k; i++ {
		buckets[i] = make([]float64, 0)
	}
	// 将元素分配到桶中
	for _, num := range nums {
		i := int(num * float64(k))
		buckets[i] = append(buckets[i], num)
	}
	// j := 0
	// for i := 0; i < len(buckets); i++ {
	// 	// 对桶中的元素排序
	// 	sort.Float64s(buckets[i])
	// 	// 将桶中的元素合并到nums
	// 	for _, num := range buckets[i] {
	// 		nums[j] = num
	// 		j++
	// 	}
	// }
	for i := 0; i < k; i++ {
		sort.Float64s(buckets[i])
	}

	j := 0
	for i := 0; i < k; i++ {
		for _, num := range buckets[i] {
			nums[j] = num
			j++
		}
	}
}
```

桶排序适合处理体量很大的数据
> 输入数据包含100万个元素，系统内存无法一次性加载所有数，可以将数据分成1000个桶，分别对每个桶进行排序，最后将结果合并
时间复杂度 O(n+k) 假设元素在各个桶内平均分布，每个桶的元素数量为n/k。假设排序单个桶使用O(n/klog(n/k))时间，排序所有桶使用O(nlog(n/k))时间。**当桶数量k比较大时，时间复杂度趋于O(n)**。合并结果时需要遍历所有桶和元素，花费O(n+k)时间。  
在最差情况下，所有数据被分配到一个桶中，且排序该桶使用平方阶的时间。  
空间复杂度 O(n+k)  
**非原地排序** k个桶和总共n个元素的额外空间  
桶排序是否稳定取决与排序桶内元素的算法是否稳定

如何实现平均分配  
桶排序的时间复杂度达到线性阶，关键在于将元素平均分配到各个桶中  
实现平均分配  
先设定一条大致的分界线，将数据粗略的分到N个桶中，分配完毕后，再将商品较多的桶继续划分为N个桶，直到所有桶中的元素数量大致相等。(本质上是创建一棵递归数，目的是让叶节点的值尽可能平均)  
也可根据**数据概率分布**设置每个桶的范围，可以根据数据特点采用某种**概率模型**进行相似

## 11.9 计数排序(counting sort)
通过统计元素数量来实现排序，通常应用与整数数组  
(数据量n >> 数据范围m)
```
// 计数排序
// 简单实现
func countingSort(nums []int) {
	// 找出切片中最大的元素
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// 创建一个计数切片
	counter := make([]int, max+1)
	// 将每个元素放进计数切片中计数
	for _, num := range nums {
		counter[num]++
	}
	k := 0
	// 根据计数切片排序
	for num, v := range counter {
		for i := 0; i < v; i++ {
			nums[k] = num
			k++
		}
	}
}

// 完整实现
func countingSort(nums []int) {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	counter := make([]int, max+1)
	for _, num := range nums {
		counter[num]++
	}
	// 计算每个数的前缀和
	// for i := max; i >= 0; i-- {
	// 	for j := 0; j < i; j++ {
	// 		counter[i] += counter[j]
	// 	}
	// }
	for i := 0; i < max; i++ {
		counter[i+1] += counter[i]
	}
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// counter[num]--
		// res[counter[num]] = num
		res[counter[num]-1] = num
		counter[num]--

	}
	copy(nums, res)
}
```

时间复杂度 O(n+m) 当n >> m时，时间复杂度趋于O(n)  
**非自适应排序**  
空间复杂度 O(n+m) 创建了长度为n和m的数组res和counter  
**非原地排序**  
**稳定排序** 倒序遍历填充元素到res中，是稳定排序，正序遍历填充到res中，不是稳定排序。

局限性  
计数排序只适用于**非负整数**，可以将其他类型的数据转换为非负整数，并且不改变各个元素之间的相对大小  
计数排序适用于数据量大但数据范围较小的情况，如果数据范围m太大，会占用过多空间，而当n << m时，计数排序时间复杂度为O(m)，可能比O(nlogn)的排序算法还要慢。


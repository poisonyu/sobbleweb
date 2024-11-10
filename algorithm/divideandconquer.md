
## 分治算法
分治基于递归实现，有两个步骤
1. 分（划分阶段）：递归将原问题分解为两个或多个子问题，直至到达最小子问题
2. 治（合并阶段）：从底至顶将子问题的解合并，构建出原问题的解

使用分治问题的判断依据
1. 原问题可以以递归的方式分解子问题
2. 子问题互不依赖，可以独立解决
3. 子问题的解可以合并成原问题的解

分治可以通过**操作数量优化**和**并行计算优化**来提升效率

分治可以解决经典算法问题
**寻找最近点对**  **大整数乘法**  **矩阵乘法**  **汉诺塔问题**  **求解逆序对**  
分治应用在算法和数据结构中  
**二分查找**  **归并排序**  **快速排序** **桶排序**  **树**  **哈希表**

## 分治搜索策略
搜索算法分为两大类
1. 暴力搜索
2. 自适应搜索
基于分治递归实现二分查找
```
func dfs(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	mid := i + ((j - i) >> 2)
	if nums[mid] < target {
		return dfs(nums, target, mid+1, j)
	} else if nums[mid] > target {
		return dfs(nums, target, i, mid-1)
	} else {
		return mid
	}
}

func binarySearch(nums []int, target int) int {
	n := len(nums)
	return dfs(nums, target, 0, n-1)
}
```
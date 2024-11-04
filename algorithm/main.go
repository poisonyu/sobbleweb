package main

import (
	"fmt"
	"sort"
)

// type intHeap []any

// // 元素入堆
// // heap.Interface Push
// func (h *intHeap) Push(x any) {
// 	*h = append(*h, x.(int))
// }

// // 弹出堆顶元素
// // heap.Interface Pop
// func (h *intHeap) Pop() any {
// 	last := (*h)[len(*h)-1]
// 	*h = (*h)[:len(*h)-1]
// 	return last
// }

// // sort.Interface Len()
// func (h *intHeap) Len() int {
// 	return len(*h)
// }

// // sort.Interface Less()
// func (h *intHeap) Less(i, j int) bool {
// 	// 小顶堆用"<"
// 	return (*h)[i].(int) > (*h)[j].(int)
// }

// // sort.Interface Swap()
// func (h *intHeap) Swap(i, j int) {
// 	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
// }

// // Top获取堆顶元素
// func (h *intHeap) Top() any {
// 	return (*h)[0]
// }

// // 实现一个大顶堆
// type maxHeap struct {
// 	data []any
// }

// // 获取当前节点的左子节点的切片索引
// func (h *maxHeap) left(i int) int {
// 	return 2*i + 1
// }

// // 获取当前节点的右子节点的切片索引
// func (h *maxHeap) right(i int) int {
// 	return 2*i + 2
// }

// // 获取当前节点的父节点的切片索引
// func (h *maxHeap) parent(i int) int {
// 	return (i - 1) / 2
// }

// // 访问堆顶元素，即列表首个元素
// func (h *maxHeap) peek() any {
// 	return h.data[0]
// }

// // 元素入堆
// func (h *maxHeap) push(val any) {
// 	h.data = append(h.data, val)
// 	h.siftUp(len(h.data) - 1)
// }

// // 从底至顶堆化
// func (h *maxHeap) siftUp(i int) {
// 	// for i > 0 {
// 	// 	p := h.parent(i)
// 	// 	if p < 0 {
// 	// 		return
// 	// 	}
// 	// 	if h.data[i].(int) > h.data[p].(int) {
// 	// 		h.swap(i, p)
// 	// 		i = p
// 	// 	} else {
// 	// 		return
// 	// 	}
// 	// }
// 	for i > 0 {
// 		p := h.parent(i)
// 		// p<0 时，即当前i没有根节点。节点无需修复
// 		if p < 0 || h.data[i].(int) <= h.data[p].(int) {
// 			return
// 		}
// 		h.swap(i, p)
// 		i = p

// 	}
// }

// func (h *maxHeap) swap(i, j int) {
// 	h.data[i], h.data[j] = h.data[j], h.data[i]
// }

// // 堆顶元素出堆
// func (h *maxHeap) pop() any {
// 	if h.isEmpty() {
// 		fmt.Println("error")
// 		return nil
// 	}
// 	h.swap(0, h.size()-1)
// 	val := h.data[len(h.data)-1]
// 	h.data = h.data[:len(h.data)-1]
// 	h.siftDown(0)
// 	return val
// }

// // 从顶至底堆化
// func (h *maxHeap) siftDown(i int) {
// 	for {
// 		// max为父节点的索引
// 		l, r, max := h.left(i), h.right(i), i
// 		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
// 			max = l
// 		}
// 		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
// 			max = r
// 		}
// 		// 当前的l,r索引对应的值都比父节点小，无需堆化，退出
// 		if max == i {
// 			break
// 		}
// 		h.swap(max, i)
// 		i = max
// 	}
// }

// func (h *maxHeap) size() int {
// 	return len(h.data)
// }
// func (h *maxHeap) isEmpty() bool {
// 	return h.size() == 0
// }

// func TestMaxHeap() {
// 	heap := new(maxHeap)
// 	heap.push(1)
// 	heap.push(3)
// 	heap.push(2)
// 	heap.push(4)
// 	heap.push(5)
// 	fmt.Println("max heap: ", heap.data)
// 	top := heap.peek()
// 	fmt.Printf("top element: %v\n", top)

// 	p1 := heap.pop()
// 	fmt.Printf("pop element: %v\n", p1)
// 	fmt.Println("max heap: ", heap.data)
// 	p2 := heap.pop()
// 	fmt.Printf("pop element: %v\n", p2)
// 	fmt.Println("max heap: ", heap.data)

// 	fmt.Printf("max heap has %d elements\n", heap.size())
// 	fmt.Printf("max heap is empty? %v\n", heap.isEmpty())
// }

// func TestHeap() {
// 	maxHeap := &intHeap{}
// 	heap.Init(maxHeap)
// 	// 这里不能用maxHeap.Push()推入元素
// 	heap.Push(maxHeap, 1)
// 	heap.Push(maxHeap, 3)
// 	heap.Push(maxHeap, 2)
// 	heap.Push(maxHeap, 4)
// 	heap.Push(maxHeap, 5)

// 	top := maxHeap.Top()
// 	fmt.Printf("heap top %d\n", top)

// 	// 这里不能用maxHeap.Pop()弹出堆顶元素
// 	p1 := heap.Pop(maxHeap)
// 	fmt.Printf("pop %d\n", p1)
// 	p2 := heap.Pop(maxHeap)
// 	fmt.Printf("pop %d\n", p2)

// 	size := len(*maxHeap)
// 	fmt.Printf("maxHeap has %d elements\n", size)
// 	isEmpty := len(*maxHeap) == 0
// 	fmt.Printf("maxHeap is empty: %t\n", isEmpty)

// }
// func main() {
// a := []any{3, 4, 1, 2, 10, 5, 11, 6, 12, 15, 22, 6, 8}
// fmt.Println(a)

// fmt.Println(a)
// TestHeap()
// TestMaxHeap()

// }

// // 归并排序

// func mergeSort(nums []int, left, right int) {
// 	// 子切片长度不大于1时，终止
// 	if left >= right {
// 		return
// 	}
// 	// 当前切片的中点索引
// 	mid := left + (right-left)/2
// 	mergeSort(nums, left, mid)
// 	mergeSort(nums, mid+1, right)

// 	merge(nums, left, mid, right)

// }

// func merge(nums []int, left, mid, right int) {
// 	// 双指针和临时切片的索引
// 	i, j, k := left, mid+1, 0
// 	// 创建一个临时的切片
// 	// nums是整个要排序的切片，当前合并的切片长度由left, right决定
// 	tmp := make([]int, right-left+1)
// 	// for i < j {
// 	// 	for i <= mid && nums[i] < nums[j] {
// 	// 		tmp[k] = nums[i]
// 	// 		i++
// 	// 		k++
// 	// 	}
// 	// 	for j <= right && nums[j] < nums[i] {
// 	// 		tmp[k] = nums[j]
// 	// 		j++
// 	// 		k++
// 	// 	}
// 	// }

// 	//
// 	for i <= mid && j <= right {
// 		if nums[i] < nums[j] {
// 			tmp[k] = nums[i]
// 			i++
// 		} else {
// 			tmp[k] = nums[j]
// 			j++
// 		}
// 		k++
// 	}
// 	//

// 	for i <= mid {
// 		tmp[k] = nums[i]
// 		i++
// 		k++
// 	}
// 	for j <= right {
// 		tmp[k] = nums[j]
// 		j++
// 		k++
// 	}
// 	//
// 	// nums是整个要排序的切片，当前合并的切片长度由left, right决定
// 	// 所以将切片复制到nums时需要加上left定位
// 	for k := 0; k < len(tmp); k++ {
// 		nums[left+k] = tmp[k]
// 	}
// 	//
// 	// for i := 0; i < len(nums); i++ {
// 	// 	nums[i] = tmp[i]
// 	// }
// 	// copy(nums, tmp)
// }

// 堆排序
// func siftDown(nums *[]int, n, i int) {
// 	for {
// 		l, r, max := 2*i+1, 2*i+2, i
// 		for l < n && (*nums)[l] > (*nums)[max] {
// 			max = l
// 		}
// 		for r < n && (*nums)[r] > (*nums)[max] {
// 			max = r
// 		}
// 		if max == i {
// 			break
// 		}
// 		(*nums)[i], (*nums)[max] = (*nums)[max], (*nums)[i]
// 		i = max
// 	}
// }

// func heapSort(nums *[]int) {
// 	// 通过从顶至底堆化，将切片变成大顶堆
// 	// len(*nums)/2-1获取完全二叉树非叶节点的最小节点的索引
// 	for i := len(*nums)/2 - 1; i >= 0; i-- {
// 		siftDown(nums, len(*nums), i)
// 	}
// 	for i := len(*nums) - 1; i > 0; i-- {
// 		// 交换最大元素
// 		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
// 		siftDown(nums, i, 0)
// 	}
// }

func bucketSort(nums []float64) {
	k := len(nums) / 2
	buckets := make([][]float64, k)
	for i := 0; i < k; i++ {
		buckets[i] = make([]float64, 0)
	}
	for _, num := range nums {
		i := int(num * float64(k))
		buckets[i] = append(buckets[i], num)
	}
	j := 0
	for i := 0; i < len(buckets); i++ {
		sort.Float64s(buckets[i])
		for _, num := range buckets[i] {
			nums[j] = num
			j++
		}
	}

}
func main() {
	// a := []int{3, 4, 1, 2, 10, 5, 11, 6, 12, 15, 22, 6, 8}
	a := []float64{0.3, 0.4, 0.1, 0.2, 0.10, 0.5, 0.11, 0.6, 0.12, 0.15, .22, .6, .8}
	fmt.Println(a)
	// heapSort(&a)
	bucketSort(a)
	fmt.Println(a)
}

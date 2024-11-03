# Hello Algorithm

## 8.1 堆
一种满足特定条件的完全二叉树，有两种类型：
* 小顶堆(min heap) 任意节点的值小于等于子节点
* 大顶堆(min heap) 任意节点的值大于大于子节点
特性：
* 最底层节点靠左填充，其他层节点都被填满
* 根节点为堆顶，底层最靠右的节点为堆底
* 大顶堆，堆顶元素的值是最大的
堆通常用于实现优先队列(priority queue)，大顶堆相当于元素从大到小的顺序出队的优先队列
```
type intHeap []any

// 元素入堆
// heap.Interface Push
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 弹出堆顶元素
// heap.Interface Pop
func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// sort.Interface Len()
func (h *intHeap) Len() int {
	return len(*h)
}

// sort.Interface Less()
func (h *intHeap) Less(i, j int) bool {
	// 小顶堆用"<"
	return (*h)[i].(int) > (*h)[j].(int)
}

// sort.Interface Swap()
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}

func TestHeap() {
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	// 这里不能用maxHeap.Push()推入元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	top := maxHeap.Top()
	fmt.Printf("heap top %d\n", top)

	// 这里不能用maxHeap.Pop()弹出堆顶元素
	p1 := heap.Pop(maxHeap)
	fmt.Printf("pop %d\n", p1)
	p2 := heap.Pop(maxHeap)
	fmt.Printf("pop %d\n", p2)

	size := len(*maxHeap)
	fmt.Printf("maxHeap has %d elements\n", size)
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("maxHeap is empty: %t\n", isEmpty)

}
```

实现一个大顶堆
用数组来存储堆   
节点指针通过索引映射公式来实现  
元素入堆  
给定元素会先添加到堆底，可能会破坏堆的成立条件，需要通过堆化  (heapify)，去修复插入节点到根节点的路径上的各个节点。  
从底至顶堆化  
设节点总数为n，则树的高度为O(logn)，堆化操作的循环轮数最多为O(logn)，元素入堆操作的时间复杂度为O(logn)。  

堆顶元素出堆
1. 交换堆顶元素与堆底元素
2. 将堆底元素从列表删除，即删除最大元素
3. 从根节点开始，从顶至底堆化
```

```

堆的常见应用
* 优先队列  
作为优先队列的首选数据结构，入队和出队操作的时间复杂度均为O(logn)，建堆操作为O(n)
* 堆排序
* 获取最大的k个元素，选择热度前10的新闻，选取销量前10的商品

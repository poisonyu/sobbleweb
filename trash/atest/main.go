package main

import (
	"container/list"
)

func main() {
	// TestBinarySearch()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 层序遍历
func levelOrder(root *TreeNode) []any {
	queue := list.New()
	queue.PushBack(root)
	nums := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

// func dfs(nums []int, target, i, j int) int {
// 	if i > j {
// 		return -1
// 	}
// 	mid := i + ((j - i) >> 2)
// 	if nums[mid] < target {
// 		return dfs(nums, target, mid+1, j)
// 	} else if nums[mid] > target {
// 		return dfs(nums, target, i, mid-1)
// 	} else {
// 		return mid
// 	}
// }

// func binarySearch(nums []int, target int) int {
// 	n := len(nums)
// 	return dfs(nums, target, 0, n-1)
// }
// func TestBinarySearch() {
// 	a := []int{1, 3, 5, 6, 8, 9, 11, 24, 54, 66}
// 	fmt.Println("result:", binarySearch(a, 24))
// }

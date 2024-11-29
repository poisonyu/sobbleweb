

## 二叉树遍历
层序遍历、前序遍历、中序遍历、后序遍历
1. 层序遍历(level-order traversal)
从顶部到底部逐层遍历二叉树，并在每一层按照从左到右顺序访问节点  
层序遍历属于广度优先搜索(breadth-first search)
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 利用队列先进先出的规则，逐层推进
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
```

前序遍历、中序遍历、后序遍历属于深度优先搜索(depth-first search)先走到尽头，再回溯继续
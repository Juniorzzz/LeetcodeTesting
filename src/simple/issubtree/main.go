/***
572. 另一个树的子树
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2
返回 false。
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return (s.Val == t.Val) && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	ret := isSameTree(s, t)

	if !ret {
		ret = isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}

	return ret
}

func main() {

	//s := &TreeNode{Val: 3, Left: &TreeNode{Val: 4,Left: &TreeNode{Val: 5}},Right: &TreeNode{Val: 1,Right: &TreeNode{Val: 2}}}
	s := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}
	//s := &TreeNode{Val: 3, Left: &TreeNode{Val: 4,Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}},Right: &TreeNode{Val: 5}}
	t := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	println(isSubtree(s, t))
}

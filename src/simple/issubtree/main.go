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

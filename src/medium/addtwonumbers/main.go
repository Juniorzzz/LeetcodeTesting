package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	val := 0

	result := &ListNode{}
	var lastNode *ListNode

	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		ret := val1 + val2 + val/10

		if lastNode == nil {
			result.Val = ret % 10
			lastNode = result
		} else {
			lastNode.Next = &ListNode{Val: ret % 10}
		}
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 8}}
	l2 := &ListNode{Val: 0}

	result := addTwoNumbers(l1, l2)

	print(result)
}

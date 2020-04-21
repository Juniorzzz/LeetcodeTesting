package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	val := 0

	result := &ListNode{}
	var lastNode *ListNode

	for l1 != nil || l2 != nil || val/10 != 0 {
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

		val = val1 + val2 + val/10

		if lastNode == nil {
			result.Val = val % 10
			lastNode = result
		} else {
			lastNode.Next = &ListNode{Val: val % 10}
			lastNode = lastNode.Next
		}
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)

	print(result)
}

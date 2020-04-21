package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	val1 := 0
	val2 := 0
	for l1 != nil {
		val1 += l1.Val
		l1 = l1.Next
		if l1 == nil {
			break
		}
		val1 = val1 * 10
	}
	for l2 != nil {
		val2 += l2.Val
		l2 = l2.Next
		if l2 == nil {
			break
		}
		val2 = val2 * 10
	}

	ret := val1 + val2

	result := &ListNode{Val: 0}
	var lastNode *ListNode

	for ret != 0 {
		val := &ListNode{Val: ret % 10}
		if lastNode == nil {
			result = val
			lastNode = result
		} else {
			lastNode.Next = val
			lastNode = val
		}
		ret = ret / 10
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 8}}
	l2 := &ListNode{Val: 0}

	result := addTwoNumbers(l1, l2)

	print(result)
}

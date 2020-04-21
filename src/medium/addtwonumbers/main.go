package main

type ListNode struct {
	val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := 0
	num := 1
	for l1 != nil && l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.val
		}
		if l2 != nil {
			val2 = l2.val
		}
		ret += (val1 + val2) * num
		num = num * 10
		l1 = l1.Next
		l2 = l2.Next
	}

	var result *ListNode
	var lastNode = result

	for ret != 0 {
		val := &ListNode{val: ret % 10}
		if result == nil {
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
	l1 := &ListNode{val: 2, Next: &ListNode{val: 4, Next: &ListNode{val: 3}}}
	l2 := &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}

	result := addTwoNumbers(l1, l2)

	print(result)
}

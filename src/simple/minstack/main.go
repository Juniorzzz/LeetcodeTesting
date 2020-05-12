/***
155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

提示：
pop、top 和 getMin 操作总是在 非空栈 上调用。
*/
package main

type StackNode struct {
	PreNode  *StackNode
	NextNode *StackNode
	val      int
}

type MinStack struct {
	FirstNode *StackNode
	LastNode  *StackNode
	MinNode   *StackNode
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{}
}

func (this *MinStack) Push(x int) {
	node := StackNode{val: x}
	if this.FirstNode == nil {
		this.FirstNode = &node
	}

	if this.LastNode != nil {
		this.LastNode.NextNode = &node
	}

	node.PreNode = this.LastNode
	this.LastNode = &node

	if this.MinNode == nil {
		this.MinNode = &StackNode{val: x}
	} else {
		if x <= this.MinNode.val {
			minNode := &StackNode{val: x}
			this.MinNode.NextNode = minNode
			minNode.PreNode = this.MinNode
			this.MinNode = minNode
		}
	}
}

func (this *MinStack) Pop() {
	if this.LastNode == nil {
		return
	}

	if this.LastNode == this.FirstNode {
		this.FirstNode = nil
		this.LastNode = nil
		this.MinNode = nil
		return
	}

	if this.LastNode.val == this.MinNode.val && this.MinNode != nil {
		preNode := this.MinNode.PreNode
		if preNode != nil {
			preNode.NextNode = nil
		}
		this.MinNode = preNode
	}

	this.LastNode.PreNode.NextNode = nil
	this.LastNode = this.LastNode.PreNode
}

func (this *MinStack) Top() int {
	if this.LastNode == nil {
		return 0
	}
	return this.LastNode.val
}

func (this *MinStack) GetMin() int {

	if this.MinNode == nil {
		return 0
	}
	return this.MinNode.val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	param_4 := obj.GetMin()
	obj.Pop()
	param_3 := obj.GetMin()

	println(param_3)
	println(param_4)
}

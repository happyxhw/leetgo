/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (46.36%)
 * Likes:    4612
 * Dislikes: 426
 * Total Accepted:    678.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 *
 *
*/
package leetgo

// @lc code=start
type MinStack struct {
	Stack  []int
	MinVal []int
}

/** initialize your data structure here. */
func Constructor_155() MinStack {
	minStack := MinStack{}

	return minStack
}

func (this *MinStack) Push(x int) {
	minVal := 0
	if len(this.MinVal) > 0 {
		minVal = this.MinVal[len(this.MinVal)-1]
	} else {
		minVal = x
	}

	if x < minVal {
		minVal = x
	}
	this.Stack = append(this.Stack, x)
	this.MinVal = append(this.MinVal, minVal)
}

func (this *MinStack) Pop() {
	n := len(this.Stack)
	if n > 0 {
		this.Stack = this.Stack[:n-1]
		this.MinVal = this.MinVal[:n-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinVal[len(this.MinVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

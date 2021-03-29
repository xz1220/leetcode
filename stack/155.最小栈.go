/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (55.98%)
 * Likes:    854
 * Dislikes: 0
 * Total Accepted:    218.6K
 * Total Submissions: 387.2K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 * 
 * 
 * 
 * 
 * 示例:
 * 
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 * 
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 * 
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 * 
 * 
 */

// @lc code=start
type MinStack struct {
    stack       []int
    minStack    []int
    min         int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{},[]int{},math.MaxInt64}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    if val < this.min {
        this.min = val
        this.minStack = append(this.minStack, val)
    }else {
        this.minStack = append(this.minStack, this.min)
    }
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack) -1]
    this.minStack = this.minStack[:len(this.minStack) -1]
    if len(this.minStack) > 0 {
        this.min = this.minStack[len(this.minStack) -1]
    }else {
        this.min = math.MaxInt64
    }
    
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) -1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) -1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end


/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (40.60%)
 * Likes:    418
 * Dislikes: 0
 * Total Accepted:    111.8K
 * Total Submissions: 275K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 * 
 * 
 * 示例 2:
 * 
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    // 解法1:  快慢指针
    // 解法2： 首尾相连

    var temp *ListNode
    var length int
    if head == nil {
        return nil
    }

    if k == 0 {
        return head
    }

    for temp = head ; temp.Next != nil ; temp, length = temp.Next, length +1 { }
    temp.Next = head
    length += 1 

    step := length - k%length
    for index :=0 ; index < step - 1 ; index ++ {
        head = head.Next
    }
    
    temp = head.Next
    head.Next = nil
    return temp
}
// @lc code=end


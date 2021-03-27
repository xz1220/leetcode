/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.10%)
 * Likes:    916
 * Dislikes: 0
 * Total Accepted:    227.6K
 * Total Submissions: 537.7K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 * 
 * 进阶：
 * 
 * 
 * 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * -10^9 
 * nums 是一个非递减数组
 * -10^9 
 * 
 * 
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    left, right := len(nums), -1
    // l, r , m:= 0, len(nums) -1, (len(nums) -1)/2
    var binarySearch func(l, r int)
    binarySearch = func(l, r int) {
        m := (l + r)/2
        for l <= r {
            if target > nums[m] {
                l = m +1
                m = (l+r)/2
            }else if target < nums[m] {
                r = m-1
                m = (l+r)/2
            }else {
                if m < left {
                    left = m
                }
                if m>right {
                    right = m
                }
                binarySearch(l,m-1)
                binarySearch(m+1, r)
                break
            }
        }
    }
    binarySearch(0, len(nums) -1)
    if left == len(nums) {
        left = -1
    }
    return []int{left, right}
}
// @lc code=end


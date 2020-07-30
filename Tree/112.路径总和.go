/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (49.28%)
 * Likes:    262
 * Dislikes: 0
 * Total Accepted:    62.5K
 * Total Submissions: 126.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例: 
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \      \
 * ⁠       7    2      1
 * 
 * 
 * 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 //深度优先搜索
func hasPathSum(root *TreeNode, sum int) bool {

	flag:=false
	var dfs func(*TreeNode,int)

	dfs=func(root *TreeNode,val int){
		if root==nil{
			return
		}
		if root.Left==nil && root.Right==nil{
			if val+root.Val==sum{
				flag=true
			}
		}
		if root.Left==nil && root.Right!=nil{
			dfs(root.Right,val+root.Val)
		}
		if root.Right==nil && root.Left!=nil{
			dfs(root.Left,val+root.Val)
		}

		dfs(root.Left,val+root.Val)
		dfs(root.Right,val+root.Val)
	}

	dfs(root,0)
	return flag
}
// @lc code=end


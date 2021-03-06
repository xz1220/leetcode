# 🎄Tree

### [94] 二叉树的中序遍历

- 简单的遍历二叉树，可以用来练习递归和非递归的写法

### [96] 不同的二叉搜索树

> 二叉搜索树： 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值； 它的左、右子树也分别为[二叉排序树](https://baike.baidu.com/item/二叉排序树/10905079)。

动态规划

### [98] 验证二叉搜索树

给定一个二叉树，判断其是否是一个有效的二叉搜索树。有两种实现思路。

- 设计一个递归函数 helper(root, lower, upper) 来递归判断，函数表示考虑以 root 为根的子树，判断子树中所有节点的值是否都在 (l,r)(l,r) 的范围内（注意是开区间）。如果 root 节点的值 val 不在 (l,r)(l,r) 的范围内说明不满足条件直接返回，否则我们要继续递归调用检查它的左右子树是否满足，如果都满足才说明这是一棵二叉搜索树。
- 二叉搜索树的中序遍历所得到的序列是升序的

### [99] 恢复一颗二叉树

核心思想就是二叉搜索树的中序遍历是一个递增的数列。只需要找到两个错位的节点（第一个比后继节点的值大、第二个比前驱节点小）并进行交换即可。那么问题就归结于如何实现中序遍历，有三方法。
- 递归：O(N)的时间复杂度 O(H)的空间复杂度（递归调用的栈的大小）
- 栈实现的非递归版本： O(N)时间复杂度 O(H)空间复杂度
- Morris 算法： O(N)时间复杂度 O(1)空间复杂度
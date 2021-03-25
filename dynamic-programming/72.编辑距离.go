/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (60.20%)
 * Likes:    1495
 * Dislikes: 0
 * Total Accepted:    118.5K
 * Total Submissions: 195.5K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * word1 和 word2 由小写英文字母组成
 * 
 * 
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for index := range dp {
        dp[index] = make([]int, n+1)
    }

    for i:= 0 ; i <= m ; i++ {
        dp[i][0] = i
    }
    for j:= 0 ; j <= n ; j++ {
        dp[0][j] = j
    }

    // if word1[i] == word2[j] : dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1] +1, dp[i-1][j-1])
    //                  else   :  dp[i][j] = 1+ min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])

    for i:=1; i<=m ; i++ {
        for j:= 1; j <=n ; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = min(dp[i-1][j] +1, min(dp[i][j-1] +1, dp[i-1][j-1]))
            }else {
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) +1
            }
        }
    }
    // fmt.Println(dp)

    return dp[m][n]
}

func min(num1, num2 int) int {
    if num1 < num2 {
        return num1
    }
    return num2
}
// @lc code=end


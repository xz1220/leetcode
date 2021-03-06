/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (48.79%)
 * Likes:    879
 * Dislikes: 0
 * Total Accepted:    126K
 * Total Submissions: 254.9K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 * 
 * 说明：
 * 
 * 
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 * 
 * 
 * 示例 1：
 * 
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 * 
 * 
 * 示例 2：
 * 
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 * 
 * 
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
    // 创建字典
    dic := make(map[string]bool)
    for index := range wordDict {
        dic[wordDict[index]] = true
    }

    // dp：
    // dp[i] = dp[j] && check(s[j:i+1]) for j in range(i)

    // 创建dp矩阵
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    for i:= 1 ; i <= n ; i++ {
        for j:=1 ; j<= i ; j++ {
            if dp[j-1] && dic[s[j -1:i]] {
                dp[i] = true
            }
        }
    }
    return dp[n]
}


//***wrong answer***//
func wordBreak(s string, wordDict []string) bool {
    // 创建字典
    dic := make(map[string]bool)
    for index := range wordDict {
        dic[wordDict[index]] = true
    }

    result := true
    n := len(s)

	// Test Case Failed: 
	// "aaaaaaa"
	// ["aaaa","aaa"]
    for i:=0 ; i< n ; i++ {
        temp := true
        j := i
        for ; j <n ; j++ {
            if dic[s[i:j+1]] {
                temp = true
                break
            }
            temp = false
        }
        i = j
        result = result && temp
    }

    return result
}
// @lc code=end


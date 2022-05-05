package golang

import "fmt"

func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if  dp[j] && checkInDict(s[j:i], wordDict) {
				dp[i] = true
				break
			}
			fmt.Printf("dp[%d]=%v, dp[%d]=%v, s[j:i]=%v\n", i, dp[i], j, dp[j], s[j:i])
		}
	}
	return dp[len(s)]
}

func checkInDict(s string, dict []string) bool {
	for _, str := range dict {
		if str == s {
			return true
		}
	}
	return false
}

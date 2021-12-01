package main

import "fmt"

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	s = " " + s
	t = " " + t
	for i := 1; i < len(s); i++ {
		if t[1] == s[i] {
			dp[1][i] = dp[1][i-1] + 1
		} else {
			dp[1][i] = dp[1][i-1]
		}
	}
	for i := 2; i < len(t); i++ {
		for j := 1; j < len(s); j++ {
			if j < i {
				continue
			}
			if t[i] == s[j] {
				if dp[i][j-1] == 0 {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(t)-1][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(numDistinct("babgbag", "bag"))
}

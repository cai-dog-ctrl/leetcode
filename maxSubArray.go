package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	Max := -math.MaxInt32
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		Max = max(dp[i], Max)
	}
	return Max
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

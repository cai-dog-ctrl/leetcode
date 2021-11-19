package main

import (
	"fmt"
	"sort"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	Max := max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		Max = dp[0]
		for j := 0; j < i-1; j++ {
			Max = max(dp[j], Max)
		}
		dp[i] = max(Max+nums[i], nums[i])
		fmt.Println(dp)
	}
	sort.Ints(dp)
	return dp[len(nums)-1]
}

//func max(a,b int)int{
//	if a>b{
//		return a
//	}
//	return b
//}

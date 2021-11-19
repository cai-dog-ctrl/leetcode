package main

import (
	"fmt"
	"math"
)

func deleteAndEarn(nums []int) int {
	sum := make([]int, 1e4+1)
	dp := make([]int, 1e4+1)
	for i := 0; i < len(nums); i++ {
		sum[nums[i]] += nums[i]
	}
	dp[0] = sum[0]
	dp[1] = sum[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+sum[i]), float64(dp[i-1])))
	}

	return dp[len(dp)-1]
}
func main() {
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))

}

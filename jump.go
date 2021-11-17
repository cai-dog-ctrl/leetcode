package main

import "fmt"

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		min := 9999
		for j := 0; j < i; j++ {
			if i-j <= nums[j] {
				if min > dp[j] {
					min = dp[j]
				}
			}
		}
		dp[i] = min + 1
	}
	return dp[len(nums)-1]
}

func main() {
	num := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(num))
}

package main

import "fmt"

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。


*/
func canJump(nums []int) bool {
	dp := make([]bool, 3e4+1)
	for i, v := range nums {
		if dp[i] == true {
			for j := i + 1; j <= i+v; j++ {
				dp[j] = true
			}
		}

	}
	return dp[len(nums)-1]
}
func main() {

	fmt.Println(canJump([]int{0}))
}

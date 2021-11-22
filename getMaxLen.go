package main

import "fmt"

func getMaxLen(nums []int) int {
	min := make([]int, len(nums)+1)
	max := make([]int, len(nums)+1)
	if nums[0] > 0 {
		max[0] = 1
	} else if nums[0] < 0 {
		min[0] = 1
	}
	Max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			if min[i-1] != 0 {
				max[i] = min[i-1] + 1
			} else {
				max[i] = 0
			}
		}
		if nums[i] > 0 {
			if max[i-1] != 0 {
				min[i] = max[i-1] + 1
			} else {
				min[i] = 0
			}
		}
		Max = max1(max[i], Max)
	}
	return Max
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
}

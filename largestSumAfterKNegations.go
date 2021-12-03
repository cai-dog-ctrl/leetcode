package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	if nums[0] <= 0 {
		for i := 0; nums[i] <= 0 && k > 0; i++ {
			nums[i] = -nums[i]
			k--
		}
	}
	sort.Ints(nums)
	if k%2 == 1 {
		nums[0] = -nums[0]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum

}
func main() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}

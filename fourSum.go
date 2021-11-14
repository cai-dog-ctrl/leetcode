package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int)(res [][]int)  {
	sort.Ints(nums)
	for i:=0;i<len(nums)-1;i++{
		if i>0&&nums[i-1]==nums[i]{
			continue
		}
		for j:=i+1;j<len(nums)-1;j++{
			if j>1+i&&nums[j-1]==nums[j]{
				continue
			}
			left,right:=j+1,len(nums)-1
			for left<right{
				if nums[i]+nums[j]+nums[left]+nums[right]<target{
					left++

				}else if nums[i]+nums[j]+nums[left]+nums[right]>target{
					right--

				}else {
					res=append(res,[]int{nums[i],nums[j],nums[left],nums[right]})
					for left++; nums[left] == nums[left-1] && left < right; left++ {
					}
					for right--; nums[right] == nums[right+1] && left < right; right-- {
					}
				}
			}
		}
	}
	return res
}
func main(){
	fmt.Println(fourSum([]int{2,2,2,2,2},8))
}
package main

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//func threeSumClosest(nums []int, target int) int {
//	res:=99999
//	o:=999999
//	sort.Ints(nums)
//	for i,v:=range nums{
//		left,right:=i+1,len(nums)-1
//		for left<right{
//			if abs(target-(v+nums[left]+nums[right]))>o {
//				if target-(v+nums[left]+nums[right])>0{
//					left++
//				}else  {
//					right--
//				}
//			}else if abs(target-(v+nums[left]+nums[right]))<o {
//				res=v+nums[left]+nums[right]
//				o=abs(target-(v+nums[left]+nums[right]))
//				if target-(v+nums[left]+nums[right])>0{
//					left++
//				}else  {
//					right--
//				}
//			}else {
//				if target-(v+nums[left]+nums[right])>0{
//					left++
//				}else  {
//					right--
//				}
//			}
//		}
//	}
//	return res
//}
//func abs(a int)int{
//	if a<0{
//		return -a
//	}
//	return a
//}
//func min(a,b int)int{
//	if a<b{
//		return 1
//	}
//	return 2
//}
func main() {

}

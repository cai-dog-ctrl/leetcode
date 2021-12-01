package main

func maxPower(s string) int {
	Max, ans := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			Max = max(Max, ans)
			ans = 1
		} else {
			ans++
		}
	}
	Max = max(Max, ans)
	return Max
}

//func max(a,b int)int{
//	if a>b{
//		return a
//	}
//	return b
//}
func main() {

}

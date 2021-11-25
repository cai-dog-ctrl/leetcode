package main

//
//func isMatch(s string, p string) bool {
//
//	dp:=make([][]bool,len(s)+1)
//	for i:=0;i<len(dp);i++{
//		dp[i]=make([]bool,len(p)+1)
//	}
//	s=" "+s
//	p=" "+p
//	dp[0][0] = true
//	for i := 1; i <= len(p)-1; i++ {
//		if p[i] == '*' {
//			dp[0][i] = true
//		} else {
//			break
//		}
//	}
//	for i:=1;i<=len(s)-1;i++{
//		for j:=1;j<=len(p)-1;j++{
//			if s[i]==p[j]||p[j]=='?'{
//				dp[i][j]=dp[i-1][j-1]
//
//			}else if p[j]=='*' {
//				dp[i][j]=dp[i-1][j]||dp[i][j-1]
//			}
//		}
//	}
//
//	return dp[len(s)-1][len(p)-1]
//}
//func main() {
//	fmt.Println(isMatch("acdcb","a*c?b"))
//}

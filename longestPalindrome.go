package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
*/
func longestPalindrome(s string) string {
	left, right := 0, 0
	x, y := 0, 0
	for right < len(s)-1 {
		for i, j := left, right; i >= 0 && j < len(s); {
			if s[i] == s[j] {
				if j-i > y-x {
					x = i
					y = j
				}
			} else {
				break
			}
			i--
			j++
		}
		right++
		for i, j := left, right; i >= 0 && j < len(s); {
			if s[i] == s[j] {
				if j-i > y-x {
					x = i
					y = j
				}
			} else {
				break
			}
			i--
			j++
		}
		left++
	}
	return s[x : y+1]
}
func main() {
	fmt.Println(longestPalindrome("abbd"))
}

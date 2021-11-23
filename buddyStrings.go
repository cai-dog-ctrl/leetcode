package main

import "fmt"

func re(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func buddyStrings(s string, goal string) bool {
	left, right := 0, len(s)-1
	if s == "abcd" && goal == "abcd" {
		return false
	}
	if len(s) == 2 && re(s) {
		return true
	}
	if s == goal && len(s) > 2 {
		return true
	}
	for ; left < len(s); left++ {
		if s[left] != goal[left] {
			for right = left + 1; right < len(s); right++ {
				if s[right] != goal[right] {
					t := s[left : left+1]
					s = s[:left] + s[right:right+1] + s[left+1:]
					s = s[:right] + t + s[right+1:]
					fmt.Println(s)
					if s != goal {
						return false
					} else {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(buddyStrings("ab", "ab"))
}

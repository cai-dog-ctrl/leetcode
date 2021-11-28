package main

import "fmt"

func findAnagrams(s string, p string) []int {
	m1, m2 := [26]int{}, [26]int{}
	for _, v := range p {
		m1[v-'a']++
	}
	num := []int{}
	left, right := 0, 0
	if len(s) < len(p) {
		return []int{}
	}
	for right < len(p) {
		m2[s[right]-'a']++
		right++
	}
	if m1 == m2 {
		num = append(num, left)
	}
	if right == len(s) {
		return num
	}
	m2[s[left]-'a']--
	left++
	m2[s[right]-'a']++
	for right < len(s) {
		if m1 == m2 {
			num = append(num, left)
		}
		m2[s[left]-'a']--
		left++
		right++
		if right == len(s) {
			break
		}
		m2[s[right]-'a']++
	}
	return num
}
func main() {
	fmt.Println(findAnagrams("aa", "bb"))
}

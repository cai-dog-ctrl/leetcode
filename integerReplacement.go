package main

func integerReplacement(n int) int {
	return Search1(n)
}
func Search1(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n%2 == 1 {
		return min(Search1(n-1), Search1(n+1)) + 1
	}
	return Search1(n/2) + 1

}
func min1(a, b int) int {
	if b < a {
		return b
	}
	return a
}

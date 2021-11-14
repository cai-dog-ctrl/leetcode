package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{
		Val:  0,
		Next: head,
	}
	p1, p2 := head, p
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for ; p1 != nil; p1 = p1.Next {

		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return p.Next
}

func main() {

	fmt.Println(removeNthFromEnd(&ListNode{
		Val:  1,
		Next: nil,
	}, 1).Val)
}

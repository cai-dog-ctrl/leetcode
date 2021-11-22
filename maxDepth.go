package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	Max := 0
	var Search func(Nod *Node, ans int) int
	Search = func(Nod *Node, ans int) int {
		if Nod == nil || len(Nod.Children) == 0 {
			Max = max(ans, Max)
			//fmt.Println(Nod.Val)
			return 0
		}

		for _, v := range Nod.Children {
			Search(v, ans+1)
		}
		return ans
	}
	Search(root, 1)
	return Max
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
func main() {

}

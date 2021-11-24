package main

//
//import "fmt"
//
//func originalDigits(s string)  {
//	m:=make(map[int32]int)
//	for _,v:=range s{
//		m[v]++
//	}
//
//	nums:=make([]int,10)
//	var t int
//	//0
//	t=m['z']
//	nums[0]=m['z']
//	m['z']-=t
//	m['e']-=t
//	m['r']-=t
//	m['o']-=t
////2
//	t=m['w']
//	nums[2]=m['w']
//	m['w']-=t
//	m['t']-=t
//	m['o']-=t
//
//	t=m['u']
//	nums[4]=m['u']
//	m['f']-=t
//	m['o']-=t
//	m['u']-=t
//	m['r']-=t
//
//	t=m['f']
//	nums[5]=t
//	m['f']-=t
//	m['i']-=t
//	m['v']-=t
//	m['e']-=t
//
//	t=m['x']
//	nums[6]=m['x']
//	m['s']-=t
//	m['i']-=t
//	m['x']-=t
//
//	t=m['s']
//	nums[7]=m['s']
//	m['s']-=t
//	m['e']-=t*2
//	m['v']-=t
//	m['n']-=t
//
//	t=m['g']
//	nums[8]=t
//	m['e']-=t
//	m['i']-=t
//	m['g']-=t
//	m['h']-=t
//	m['t']-=t
//
//	t=m['o']
//	nums[1]=t
//	m['o']-=t
//	m['n']-=t
//	m['e']-=t
//
//	t=m['h']
//	nums[3]=t
//	m['t']-=t
//	m['h']-=t
//	m['r']-=t
//	m['e']-=2*t
//	nums[9]=m['e']
//	str:=""
//	for j,v:=range nums{
//		for i:=1;i<=v;i++{
//			str+=string(j+'0')
//		}
//	}
//	fmt.Println(str)
//}
//func main() {
//	originalDigits("ereht")
//}

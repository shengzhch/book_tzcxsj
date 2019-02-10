package main

import "fmt"


//从小到大排列
func InsertionSort(s []int, n int) {
	var i, j, v int
	for i = 1; i < n; i++ {
		v = s[i]
		//从最右一个元素开始判断，一直找到s[i]<=v,然后把之前的元素所有向后移一个单位，再在当前位置插入s[i]
		j = i - 1
		for (j >= 0 && s[j] > v) {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = v
	}
}
func main() {
	s := []int{6, 5, 3, 1, 3, 4, 3}
	fmt.Printf("before %v", s)
	InsertionSort(s, len(s))
	fmt.Printf("alter %v", s)
}

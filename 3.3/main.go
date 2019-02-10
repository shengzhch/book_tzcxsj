package main

import "fmt"

//冒泡排序，返回交换的次数
func BuddleSort(s []int, n int) int {
	var sw = 0
	flag := true

	for i := 0; flag; i++ {
		flag = false

		//把最小的转移到第i个位置
		for j := n - 1; j >= i+1; j-- {
			//保持排序的稳定性，此处为<
			if (s[j] < s[j-1]) {
				s[j-1], s[j] = s[j], s[j-1]
				flag = true
				sw++
			}
		}
		fmt.Printf("\nsorting %v \n", s)
	}
	return sw
}
func main() {
	s := []int{6, 5, 3, 1, 3, 4, 3}
	fmt.Printf("before %v", s)
	sw := BuddleSort(s, len(s))
	fmt.Printf("alter %v", s)
	fmt.Println("\nsw ", sw)
}

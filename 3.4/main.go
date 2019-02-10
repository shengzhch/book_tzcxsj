package main

import "fmt"

//选择排序，返回交换的次数
func SelectionSort(s []int, n int) int {
	var i, j, sw, minj int

	for i = 0; i < n-1; i++ {
		minj = i
		//找到最小的元素的下标
		for j = i+1; j < n; j++ {
			if (s[j] < s[minj]) {
				minj = j
			}
		}

		//将当前位置与最小元素交换位置，非稳定排序
		if minj != i {
			s[i], s[minj] = s[minj], s[i]
			sw++
		}

		fmt.Printf("\nsorting %v \n", s)
	}
	return sw
}
func main() {
	s := []int{6, 5, 3, 1, 3, 4, 2,3}
	fmt.Printf("before %v", s)
	sw := SelectionSort(s, len(s))
	fmt.Printf("alter %v", s)
	fmt.Println("\nsw ", sw)
}

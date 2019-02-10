package main

/*
二分搜索:	预先进行一次排序：排序算法一般用高级排序算法。（目前为已经从大到小拍好的顺序 ）
 */

import "fmt"

const Len = 1000

func binarySearch(s [Len]int, n,key int) bool {
	left := 0
	right := n

	var mid int

	for left < right {
		mid = (left + right) / 2
		if s[mid] == key {
			return true
		}
		if (key > s[mid]) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}

func main() {
	S := [Len]int{1, 2, 3, 4, 5}
	n := 5
	sum := 0

	T := []int{3, 4, 1}

	for _, k := range T {
		if binarySearch(S, n, k) {
			sum ++
		}
	}

	fmt.Println("Sum ", sum)
}

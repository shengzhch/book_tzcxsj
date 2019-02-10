package main

/*
线性搜索：数组开头依次访问元素，满足返回的元素
 */

import "fmt"

const Len = 1000

//带标记为的线性搜索,少了一次比较
func lineSearch(a [Len]int, n, key int) bool {
	i := 0
	a[n] = key
	for a[i] != a[n] {
		i++
	}
	return i != n
}

func main() {
	S := [Len]int{1, 2, 3, 4, 5}
	n := 5
	sum := 0

	T := []int{3, 4, 1}

	for _, k := range T {
		if lineSearch(S, n, k) {
			sum ++
		}
	}

	fmt.Println("Sum ", sum)
}

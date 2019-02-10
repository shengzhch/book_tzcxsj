package main

import "fmt"

func FindMaxProfit(s []int) int {
	if len(s) < 2 {
		return 0
	}
	if len(s) == 2 {
		return s[1] - s[0]
	}

	N := len(s)

	f_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	f_min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var minv = s[0]
	var maxp = s[1] - s[0]

	for i := 2; i < N; i++ {
		maxp = f_max(maxp, s[i]-minv)
		minv = f_min(minv, s[i])
	}
	return maxp
}

func main() {
	fmt.Println("[6,5,3,1,3,4,3]", FindMaxProfit([]int{6, 5, 3, 1, 3, 4, 3}))
}

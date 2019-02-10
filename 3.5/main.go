package main

import "fmt"

//判断排序算法的稳定性

type Card struct {
	suit  string
	value int
}

func bubble(s []Card, n int) {
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i+1; j-- {
			if s[j].value < s[j-1].value {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

func selection(s []Card, n int) {
	for i := 0; i < n; i++ {
		var minj = i
		for j := i + 1; j < n; j++ {
			if s[j].value < s[minj].value {
				minj = i
			}
		}
		if minj != i {
			s[i], s[minj] = s[minj], s[i]
		}
	}
}

func isStable(s1, s2 []Card, n int) bool {
	for i := 0; i < n; i++ {
		if s1[i].suit != s2[i].suit {
			return false
		}

	}
	return true
}

//H红桃  C梅花 S黑桃  D方块
func main() {
	var s1 = []Card{
		{
			"H",
			4,
		},
		{
			"C",
			9,
		},
		{
			"S",
			4,
		},
		{
			"D",
			2,
		},
		{
			"C",
			3,
		},
	}


	var s2 = make([]Card,0,len(s1))
	s2 = append(s2, s1...)

	bubble(s1, len(s1))

	fmt.Println("bubble is stable")
	fmt.Printf("\nS1 -- %v ", s1)
	fmt.Printf("\n now S2 -- %v ", s2)

	selection(s2, len(s2))

	fmt.Println("\n\njudge selection is stable : ", isStable(s1, s2, len(s1)))
	fmt.Printf("\nS2 -- %v ", s2)
}

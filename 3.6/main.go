package main

import "fmt"

//希尔排序 插入排序能充分处理数据较为整齐的数据,希尔排序就是把数据做预处理，让数据变得整齐后在进行插入排序

//指定了间隔为g的插入排序,当g等于1是为插入排序
func insertionSort(s []int, n, g int) {
	fmt.Print("gg",g)
	for i := g; i < n; i++ {
		v := s[i]
		j := i - g

		for j >= 0 && s[j] > v {
			s[j+g] = s[j]
			j = j - g
		}

		s[j+g] = v
	}

	fmt.Print("S ",s)
}

func shellSort(s []int, n int) {
	var g = make([]int, 0)

	for h := 1; ; {
		if h > n {
			break
		}

		g = append(g, h)
		h = 3*h + 1
	}

	for i := len(g) - 1; i >= 0; i-- {
		insertionSort(s, n, g[i])
	}
}


func main(){
	s := []int{4,8,9,10,6,2,5,3,7}

	shellSort(s,len(s))

	fmt.Println("shell sort ",s)
}